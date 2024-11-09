package controllers

import (
	"fmt"
	"go-kozel/internal/game"
	"go-kozel/internal/services"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type LobbyController struct {
	app         *pocketbase.PocketBase
	hubService  *game.LobbyHub
	authService *services.AuthService
}

func NewLobbyController(app *pocketbase.PocketBase) *LobbyController {
	return &LobbyController{
		app:         app,
		authService: services.NewAuthService(app),
	}
}

func (l *LobbyController) Register(hubService *game.LobbyHub) {
	l.hubService = hubService

	l.app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/lobby/list", l.getLobbyList)
		return nil
	})
	l.app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/lobby/:lobby_id", l.getLobby)
		return nil
	})
	l.app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/lobby/new", l.createLobby)
		return nil
	})
	l.app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/lobby/join/:lobby_id/:token", l.joinLobby)
		return nil
	})
}

type LobbyRes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Get lobby list
func (l *LobbyController) getLobbyList(c echo.Context) error {
	lobbies := make([]LobbyRes, 0)

	for _, l := range l.hubService.Lobbies {
		lobbies = append(lobbies, LobbyRes{
			Id:   l.Id,
			Name: l.Lobby.Name,
		})
	}

	c.JSON(http.StatusOK, lobbies)
	return nil
}

// Create new lobby
func (l *LobbyController) createLobby(c echo.Context) error {
	// Check auth
	info := apis.RequestInfo(c)
	user := info.AuthRecord
	if user == nil {
		err := fmt.Errorf("user not found")
		return apis.NewUnauthorizedError(err.Error(), nil)
	}
	userId := user.Id
	username := user.Username()
	fmt.Printf("Creating lobby for user %s %s\n", userId, username)
	// Проверка на существование лобби
	if l.hubService.HasLobby(userId) {
		err := fmt.Errorf("lobby already exists")
		return apis.NewBadRequestError(err.Error(), nil)
	}

	// Создаём лобби в хабе
	createdLobby := l.hubService.CreateNewLobby(userId, username)
	fmt.Printf("Lobby %s created\n", createdLobby.Id)

	c.JSON(http.StatusCreated, createdLobby.Id)

	return nil
}

// Join lobby
func (l *LobbyController) joinLobby(c echo.Context) error {
	lobbyId := c.PathParam("lobby_id")
	token := c.PathParam("token")

	// Check auth
	user, err := l.authService.GetUserFromToken(token)
	if err != nil {
		return apis.NewUnauthorizedError(err.Error(), nil)
	}

	lobbyHub := l.hubService.GetLobbyService(lobbyId)
	if lobbyHub == nil {
		err := fmt.Errorf("lobby not found")
		return apis.NewNotFoundError(err.Error(), nil)
	}

	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return apis.NewBadRequestError(err.Error(), nil)
	}

	client := game.NewWsClient(lobbyHub, &user, conn)

	if err := lobbyHub.AddClient(client); err != nil {
		return apis.NewBadRequestError(err.Error(), nil)
	}

	defer func() {
		lobbyHub.RemoveClient(client)
	}()

	// Запускаем WS слушатели
	go client.WriteMessage()
	client.ReadMessage()

	return nil
}

type LobbyDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (l *LobbyController) getLobby(c echo.Context) error {
	lobbyId := c.PathParam("lobby_id")
	lobby := l.hubService.GetLobbyService(lobbyId)

	if lobby == nil {
		err := fmt.Errorf("lobby not found")
		return apis.NewNotFoundError(err.Error(), nil)
	}

	c.JSON(http.StatusOK, LobbyDto{Id: lobby.Id, Name: lobby.Lobby.Name})

	return nil
}
