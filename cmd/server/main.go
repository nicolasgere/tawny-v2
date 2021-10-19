package main

import (
	"fmt"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"tawny/lib/front"
	pubsub "tawny/protos"
)

func main() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pubsub.RegisterPubSubServiceServer(grpcServer, &front.PushServer{})
	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Recoverer,
	)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8080"))
	if err != nil {
		grpclog.Fatalf("failed starting grpc server: %v", err)
	}
	grpcServer.Serve(lis)
}
