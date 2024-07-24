package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"ocr-service-dev/internal/handlers"
	pb "ocr-service-dev/internal/proto"
	"ocr-service-dev/internal/utils"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port           = flag.Int("port", 50051, "The server port")
	textractClient *textract.Client
)

func RunServer() {
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	textractClient = textract.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	// Connect to RDS Postgres
	db, err := utils.InitializePostgresConnection()
	if err != nil {
		log.Fatalf("unable to initialize connection, %v", err)
	}
	defer db.Close()
	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}
	log.Printf("version=%s\n", version)

	// gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOcrServiceServer(s, &handlers.OcrServiceHandler{DB: db, TextractClient: textractClient, S3Client: s3Client})
	reflection.Register(s)

	go func() {
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// gRPC-Web server
	grpcWebServer := grpcweb.WrapServer(
		s,
		// Enable CORS
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)
	srv := &http.Server{
		Handler: grpcWebServer,
		Addr:    fmt.Sprintf("localhost:%d", *port+1),
	}

	go func() {
		log.Printf("HTTP server listening at %v", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Prevent main from exiting
	select {}
}
