package main

import (
	"go-kozel/internal/services"
	router "go-kozel/internal/transport/http"
)

// func startGame() {
// 	game := domain.NewGame()
// 	game.Start()
// }

// func exitGame() {
// 	fmt.Println("Exiting game...")
// 	os.Exit(0)
// }

func main() {

	// hub := ws.NewHub()
	// wsHandler := ws.NewHandler(hub)
	// go hub.Run()

	hub := services.NewHub()

	router.InitRouter(hub)
	router.Start(":8080")
}
