package handlers

import (
	"go-kozel/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	hub *services.Hub
}

func NewHandler(h *services.Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type NewLobbyDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type LobbyRes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) NewLobby(c *gin.Context) {
	var req NewLobbyDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hasLobby := h.hub.HasLobby(req.Id)
	if hasLobby {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lobby already exists"})
		return
	}

	ownerId := req.Id
	lobbyName := req.Name
	lobby := services.NewLobby(ownerId, lobbyName)
	h.hub.AddLobby(lobby)

	c.JSON(http.StatusCreated, gin.H{"lobbyId": lobby.Id})
}

func (h *Handler) GetLobbies(c *gin.Context) {
	lobbies := make([]LobbyRes, 0)

	for _, l := range h.hub.Lobbies {
		lobbies = append(lobbies, LobbyRes{
			Id:   l.Id,
			Name: l.Name,
		})
	}

	c.JSON(http.StatusOK, lobbies)
}

func (h *Handler) WatchLobbies(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	flusher, _ := c.Writer.(http.Flusher)

	hubChan := make(chan *services.HubEvent)
	h.hub.Listen(&hubChan)
	defer func() {
		h.hub.Unlisten(&hubChan)
	}()

	for {
		if hubChan == nil {
			break
		}
		msg := <-hubChan
		if msg == nil {
			break
		}
		if msg.Event == "new_lobby" {
			c.SSEvent("new_lobby", NewLobbyDto{Id: msg.LobbyId, Name: msg.Data.Name})
		}
		if msg.Event == "remove_lobby" {
			c.SSEvent("remove_lobby", msg.LobbyId)
		}
		flusher.Flush()
	}
}
