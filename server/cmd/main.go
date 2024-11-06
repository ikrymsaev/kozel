package main

import (
	"go-kozel/internal/infrastructure/pb"
	"go-kozel/internal/services"
	router "go-kozel/internal/transport/http"
)

func main() {
	go pb.Run()

	hub := services.NewHub()

	router.InitRouter(hub)
	router.Start(":8080")
}
