package controllers

import (
	"go-kozel/internal/hub"
	"go-kozel/internal/services"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type HubController struct {
	app         *pocketbase.PocketBase
	hub         *hub.Hub
	authService *services.AuthService
}

func NewHubController(app *pocketbase.PocketBase) *HubController {
	return &HubController{
		app:         app,
		authService: services.NewAuthService(app),
	}
}

func (h *HubController) Register(hub *hub.Hub) {
	h.hub = hub

	h.app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/hub/connect/:token", h.connect)
		return nil
	})
}

func (h *HubController) connect(ctx echo.Context) error {
	token := ctx.PathParam("token")
	// Check auth
	user, err := h.authService.GetUserFromToken(token)
	if err != nil {
		return apis.NewUnauthorizedError(err.Error(), nil)
	}
	if err != nil {
		return apis.NewUnauthorizedError(err.Error(), nil)
	}

	conn, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
	if err != nil {
		return apis.NewBadRequestError(err.Error(), nil)
	}

	h.hub.Connect(conn, &user)

	defer func() {
		h.hub.Disconnect(conn)
	}()

	return nil
}
