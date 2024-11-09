package controllers

import (
	"go-kozel/internal/domain"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type SettingsController struct {
	app *pocketbase.PocketBase
}

func NewSettingsController(app *pocketbase.PocketBase) *SettingsController {
	return &SettingsController{app: app}
}

func (s *SettingsController) Register() {
	s.app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/settings/deck", s.getDeck)
		return nil
	})
}

type DeckCardRes struct {
	CardType domain.CardType `json:"suit"`
	CardSuit domain.CardSuit `json:"type"`
	IsTrump  bool            `json:"isTrump"`
	ImageUri string          `json:"imageUri"`
}

func (m *SettingsController) getDeck(c echo.Context) error {
	deck := domain.NewDeck()
	response := make([]DeckCardRes, 0)
	for _, card := range deck.Cards {
		response = append(response, DeckCardRes{
			CardType: card.CardType,
			CardSuit: card.CardSuit,
			IsTrump:  card.IsTrump,
			ImageUri: card.ImageUri,
		})
	}
	c.JSON(200, response)

	return nil
}
