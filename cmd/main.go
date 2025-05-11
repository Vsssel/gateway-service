package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "gateway-service/protogen/golang/myservice"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC handler: %v", err)
	}

	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./protogen/golang/swagger/your_service.swagger.json")
	})

	fs := http.FileServer(http.Dir("./swagger-ui"))
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	http.Handle("/", mux)

	log.Println("HTTP Gateway запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
