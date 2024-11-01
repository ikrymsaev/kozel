package handlers

import (
	"fmt"
	"go-kozel/internal/domain"
	"go-kozel/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type LobbyHandler struct {
	hub *services.Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewLobbyHandler(h *services.Hub) *LobbyHandler {
	return &LobbyHandler{
		hub: h,
	}
}

func (h *LobbyHandler) NewLobby(c *gin.Context) {
	fmt.Println(">>>NewLobby")

	userId := c.Query("user_id")
	username := c.Query("username")

	// Проверка на существование лобби
	if h.hub.HasLobby(userId) {
		fmt.Println("Lobby already exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lobby already exists"})
		return
	}

	// Создаём лобби в хабе
	createdLobby := h.hub.CreateNewLobby(userId, username) // TODO: Clear empty lobbies
	fmt.Printf("Lobby %s created\n", createdLobby.Id)

	c.JSON(http.StatusCreated, createdLobby.Id)
}

func (h *LobbyHandler) JoinLobby(c *gin.Context) {
	lobbyId := c.Param("lobby_id")
	userId := c.Query("user_id")
	username := c.Query("username")

	lobbyHub := h.hub.GetHubLobby(lobbyId)
	if lobbyHub == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "lobby not found"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Joining lobby %s\n", lobbyId)
	fmt.Printf("User %s\n", userId)

	user := &domain.User{ID: userId, Username: username}
	client := services.NewClient(lobbyHub, user, conn)

	if err := lobbyHub.AddClient(client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer func() {
		lobbyHub.RemoveClient(client)
	}()

	// Запускаем WS слушатели
	go client.WriteMessage()
	client.ReadMessage()
}
