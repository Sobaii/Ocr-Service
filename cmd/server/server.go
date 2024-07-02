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

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func RunServer() {
	flag.Parse()

	// gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOcrServiceServer(s, &handlers.OcrServiceHandler{})
	reflection.Register(s)

	var context = context.Background()
	utils.InitializeSearchIndex(context)

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
