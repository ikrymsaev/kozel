package domain

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	Cards [32]Card
}

func NewDeck() Deck {
	var cards [32]Card
	i := 0
	for _, suit := range CARD_SUITS {
		for _, cardType := range CARD_TYPES {
			uri := fmt.Sprintf("images/%d.jpg", i+1)
			cards[i] = NewCard(cardType, suit, uri)
			i++
		}
	}
	return Deck{cards}
}

// Перемашать колоду
func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// Установить козырь
func (d *Deck) SetTrump(suit *ESuit) {
	for _, card := range d.Cards {
		if card.CardSuit.Suit == *suit || card.CardType.Type == Jack {
			card.SetTrump()
		}
	}
}

func (d *Deck) CardsString() string {
	var cardsString string
	for _, card := range d.Cards {
		cardsString += card.String() + " "
	}
	return cardsString
}

func (d *Deck) GetJacksInGame() []*Card {
	var jacks []*Card
	for _, card := range d.Cards {
		if card.CardType.Type == Jack && !card.IsUsed {
			jacks = append(jacks, &card)
		}
	}
	return jacks
}

func (d *Deck) HasJackInGame() bool {
	for _, card := range d.Cards {
		if card.CardType.Type == Jack && !card.IsUsed {
			return true
		}
	}
	return false
}

func (d *Deck) GetOlderTrumpInGame() *Card {
	var olderTrump *Card
	for _, card := range d.Cards {
		if card.IsTrump && card.CardType.Type != Jack && !card.IsUsed {
			if olderTrump == nil || card.CardSuit.Order > olderTrump.CardSuit.Order {
				olderTrump = &card
			}
		}
	}
	return olderTrump
}

func (d *Deck) GetTrumpsInGame() []*Card {
	trumps := make([]*Card, 0)

	for _, card := range d.Cards {
		if card.IsTrump && card.CardType.Type != Jack && !card.IsUsed {
			trumps = append(trumps, &card)
		}
	}
	return trumps
}

func (d *Deck) GetSuitsInGame(suit *ESuit) []*Card {
	cards := make([]*Card, 0)
	for _, card := range d.Cards {
		if card.CardType.Type != Jack && !card.IsUsed && card.CardSuit.Suit == *suit {
			cards = append(cards, &card)
		}
	}
	return cards
}

func (d *Deck) IsHasAce(suit *ESuit) bool {
	for _, card := range d.Cards {
		if card.CardType.Type == Ace && card.CardSuit.Suit == *suit && !card.IsUsed {
			return true
		}
	}
	return false
}
