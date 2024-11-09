package pb

import (
	"go-kozel/internal/game"
	"go-kozel/internal/infrastructure/pb/controllers"

	"github.com/pocketbase/pocketbase"
)

type App struct {
	Instance *pocketbase.PocketBase
}

func New() *App {
	return &App{
		Instance: pocketbase.New(),
	}
}

func (p *App) Run(hubService *game.LobbyHub) {
	hub := hubService.GetHub()

	controllers.NewHubController(p.Instance).Register(hub)
	controllers.NewAuthController(p.Instance).Register()
	controllers.NewSettingsController(p.Instance).Register()
	controllers.NewLobbyController(p.Instance).Register(hubService)
	p.Instance.Start()
}
