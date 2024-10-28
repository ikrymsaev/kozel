package handlers

import (
	"go-kozel/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HubHandler struct {
	hub *services.Hub
}

func NewHubHandler(h *services.Hub) *HubHandler {
	return &HubHandler{
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

func (h *HubHandler) GetLobbies(c *gin.Context) {
	lobbies := make([]LobbyRes, 0)

	for _, l := range h.hub.Lobbies {
		lobbies = append(lobbies, LobbyRes{
			Id:   l.Id,
			Name: l.Lobby.Name,
		})
	}

	c.JSON(http.StatusOK, lobbies)
}

func (h *HubHandler) WatchLobbies(c *gin.Context) {
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
