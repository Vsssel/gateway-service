package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/third_party/googleapis/google/api/annotations"
	"google.golang.org/grpc"
	pb "your_project_path/protogen/golang/your_service"
)

func main() {
	// gRPC сервер
	grpcServer := grpc.NewServer()
	pb.RegisterYourServiceServer(grpcServer, &server{})

	// Настройка HTTP Gateway
	mux := runtime.NewServeMux()
	err := pb.RegisterYourServiceHandlerServer(context.Background(), mux, &server{})
	if err != nil {
		log.Fatalf("could not register service handler: %v", err)
	}

	// Запуск HTTP сервера
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Запуск gRPC сервера
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Запуск HTTP сервера
	log.Println("HTTP server started at :8080")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
