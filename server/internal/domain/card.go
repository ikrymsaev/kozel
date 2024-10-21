package domain

import "fmt"

type Card struct {
	CardType CardType
	CardSuit CardSuit
	IsUsed   bool
	IsTrump  bool
	Owner    *Player
}

func NewCard(cardType CardType, cardSuit CardSuit) Card {
	return Card{
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
