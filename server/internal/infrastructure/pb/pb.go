package pb

import (
	"go-kozel/internal/game"
	"go-kozel/internal/infrastructure/pb/controllers"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
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
	p.Instance.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("../static"), false))
		return nil
	})
	p.Instance.Start()
}
