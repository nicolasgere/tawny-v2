package main

import (
	"fmt"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"tawny/lib/configuration"
	"tawny/lib/front"
	"tawny/lib/logging"
	"tawny/lib/object"
	pubsub "tawny/protos"
)

func main() {
	config := configuration.InitConfig()
	logger := logging.InitLogger(config)
	logger.Info("Initializing server...")
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pubsub.RegisterPubSubServiceServer(grpcServer, &front.PushServer{
		Logger:        logger,
		Configuration: &config,
		ObjectService: object.NewObjectService(logger, config),
	})
	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Recoverer,
	)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logger.Fatal("Can't start server", zap.Error(err))
	}
	logger.Info(fmt.Sprintf("Listening %s", config.Port))
	grpcServer.Serve(lis)
}
