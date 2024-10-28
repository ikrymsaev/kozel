package main

import (
	"go-kozel/internal/services"
	router "go-kozel/internal/transport/http"
)

func main() {
	hub := services.NewHub()

	router.InitRouter(hub)
	router.Start(":8080")
}
