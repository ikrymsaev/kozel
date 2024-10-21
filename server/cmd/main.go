package main

import (
	router "go-kozel/internal/transport/http"
	"go-kozel/internal/transport/ws"
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

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(wsHandler)
	router.Start(":8080")

	// prompt := promptui.Select{
	// 	Label: "Kozel game",
	// 	Items: []string{"Start", "Exit"},
	// }

	// _, result, err := prompt.Run()

	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }

	// switch result {
	// case "Start":
	// 	startGame()
	// case "Exit":
	// 	exitGame()
	// default:
	// 	fmt.Println("Invalid choice. Exiting...")
	// 	exitGame()
	// }
}
