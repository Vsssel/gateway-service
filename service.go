package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    pb "gateway-service/protogen/golang/myservice"  // Путь к сгенерированным файлам
)

type server struct {
    pb.UnimplementedYourServiceServer
}

// Реализуем метод из сервиса (например, Echo)
func (s *server) Echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {
    return &pb.StringMessage{Value: "Hello, " + req.GetValue()}, nil
}

func Service() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterYourServiceServer(grpcServer, &server{})
    fmt.Println("Server started at :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
