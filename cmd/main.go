package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	usersvc "gateway/protogen/golang/user_service"
	ordersvc "gateway/protogen/golang/order_service"
	product "gateway/protogen/golang/product_service"
	_"gateway/middleware"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := usersvc.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "46.101.215.10:7008", opts); err != nil {
		log.Fatalf("Failed to register UserService gRPC handler: %v", err)
	}

	if err := ordersvc.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, "localhost:50053", opts); err != nil {
		log.Fatalf("Failed to register OrderService gRPC handler: %v", err)
	}

	if err := ordersvc.RegisterResourceServiceHandlerFromEndpoint(ctx, mux, "localhost:50053", opts); err != nil {
		log.Fatalf("Failed to register OrderService gRPC handler: %v", err)
	}

	if err := product.RegisterProductCatalogServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts); err != nil {
		log.Fatalf("Failed to register OrderService gRPC handler: %v", err)
	}

	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))

	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./protogen/golang/swagger/swagger.swagger.json")
	})

	http.Handle("/", mux)

	log.Println("HTTP Gateway запущен на :2222")
	log.Fatal(http.ListenAndServe(":2222", nil))
}
