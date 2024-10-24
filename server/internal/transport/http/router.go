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

	r.Static("/images", "../static/images")
	hubHandlers := handlers.NewHubHandler(hub)
	r.POST("/hub/new_lobby", hubHandlers.NewLobby)
	r.GET("/hub/lobbies", hubHandlers.GetLobbies)
	r.GET("/hub/watch_lobbies", hubHandlers.WatchLobbies)

	settingsHandlers := handlers.NewSettingsHandler()
	r.GET("/settings/deck", settingsHandlers.GetDeck)
}

func Start(addr string) error {
	return r.Run(addr)
}
