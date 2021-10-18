package server

import (
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
}
