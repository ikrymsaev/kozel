package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Card struct {
	Id       string
	ImageUri string
	CardType CardType
	CardSuit CardSuit
	IsUsed   bool
	IsTrump  bool
	Owner    *Player
}

func NewCard(cardType CardType, cardSuit CardSuit, imageUri string) Card {

	// NewCard creates a new Card instance with the given card type, card suit, and image URI.
	// The card is initially marked as not used and is considered a trump card if its type is Jack.
	return Card{
		Id:       uuid.New().String(),
		ImageUri: imageUri,
		CardType: cardType,
		CardSuit: cardSuit,
		IsUsed:   false,
		IsTrump:  cardType.Type == Jack,
	}
}

func (c *Card) SetOwner(owner *Player) {
	c.Owner = owner
}

func (c *Card) SetTrump() {
	c.IsTrump = true
}

func (c *Card) SetUsed() {
	c.IsUsed = true
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.CardType.Name, c.CardSuit.Name)
}
