package main

import (
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/api"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/api/handlers"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/clients"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	grpcClients,err:=clients.NewGrpcClients(cfg)
	if err!=nil{
		panic(err)
	}
	h := handlers.NewHandler(grpcClients,cfg)
	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	if cfg.Environment != "production" {
		router.Use(gin.Logger(), gin.Recovery())
	}

	api.SetUpApi(router, h, cfg)

	router.Run(cfg.HTTPPort)
}
