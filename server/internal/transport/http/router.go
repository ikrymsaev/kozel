package router

import (
	"go-kozel/internal/services"
	"go-kozel/internal/transport/http/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(hub *services.Hub) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	lobbyHnd := handlers.NewLobbyHandler(hub)
	r.GET("/lobby/new", lobbyHnd.NewLobby)
	r.GET("/lobby/join/:lobby_id", lobbyHnd.JoinLobby)

	r.Static("/images", "../static/images")
	hubHnd := handlers.NewHubHandler(hub)
	r.GET("/hub/lobbies", hubHnd.GetLobbies)
	r.GET("/hub/watch_lobbies", hubHnd.WatchLobbies)

	settingsHnd := handlers.NewSettingsHandler()
	r.GET("/settings/deck", settingsHnd.GetDeck)
}

func Start(addr string) error {
	return r.Run(addr)
}
