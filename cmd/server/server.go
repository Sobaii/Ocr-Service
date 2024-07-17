package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"ocr-service-dev/internal/handlers"
	pb "ocr-service-dev/internal/proto"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type ServerInterface interface {
	Serve(net.Listener) error
}

func launchServer(server ServerInterface, port int) {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		os.Exit(1)
	}
	err = server.Serve(conn)

	if err != nil {
		fmt.Printf("Serve on port %d: %v\n", port, err)
		os.Exit(1)
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for the preflight request
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Set CORS headers for the actual request
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			// forward request to gRPC server
			grpcServer.ServeHTTP(w, r)
		} else {
			// ensure application/json content type is accepted
			if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
				http.Error(w, "Invalid content type. Only application/json is accepted", http.StatusUnsupportedMediaType)
				return
			}

			// forward request to HTTP server
			otherHandler.ServeHTTP(w, r)
		}
	})
}

var (
	// port           = flag.Int("port", 50051, "The server port")
	textractClient *textract.Client
	s3Client       *s3.Client
	grpcServer     *grpc.Server
	mux            *http.ServeMux
	gwmux          *runtime.ServeMux
)

func RunServer() {

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	textractClient = textract.NewFromConfig(cfg)
	s3Client = s3.NewFromConfig(cfg)

	grpcServer = grpc.NewServer()
	mux = http.NewServeMux()
	pb.RegisterOcrServiceServer(grpcServer, &handlers.OcrServiceHandler{TextractClient: textractClient, S3Client: s3Client})
	reflection.Register(grpcServer)
	ctx := context.Background()
	gwmux = runtime.NewServeMux()

	err = pb.RegisterOcrServiceHandlerFromEndpoint(ctx, gwmux, "localhost:9090", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	mux.Handle("/", gwmux)
	// launch grpc server
	go launchServer(grpcServer, 9090)

	// launch http server redirecting to grpc
	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: grpcHandlerFunc(grpcServer, mux),
	}
	go launchServer(srv, 8080)

	select {}
}
