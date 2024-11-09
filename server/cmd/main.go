package main

import (
	"go-kozel/internal/game"
	"go-kozel/internal/hub"
	"go-kozel/internal/infrastructure/pb"
)

func main() {
	hub := hub.NewHub()
	go hub.Run()
	hubService := game.NewLobbyHub(hub)

	app := pb.New()
	app.Run(hubService)
}
