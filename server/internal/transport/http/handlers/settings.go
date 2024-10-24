package handlers

import (
	"go-kozel/internal/domain"

	"github.com/gin-gonic/gin"
)

type SettingsHandler struct{}

func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{}
}

type DeckCardRes struct {
	CardType domain.CardType `json:"suit"`
	CardSuit domain.CardSuit `json:"type"`
	IsTrump  bool            `json:"isTrump"`
}

func (h *SettingsHandler) GetDeck(c *gin.Context) {
	deck := domain.NewDeck()
	response := make([]DeckCardRes, 0)
	for _, card := range deck.Cards {
		response = append(response, DeckCardRes{
			CardType: card.CardType,
			CardSuit: card.CardSuit,
			IsTrump:  card.IsTrump,
		})
	}

	c.JSON(200, response)
}
