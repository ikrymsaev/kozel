package main

import (
	"go-kozel/internal/services"
	router "go-kozel/internal/transport/http"

	"github.com/pocketbase/pocketbase"
)

func main() {
	pb := pocketbase.New()
	go pb.Start()
	hub := services.NewHub()

	router.InitRouter(hub)
	router.Start(":8080")
}
