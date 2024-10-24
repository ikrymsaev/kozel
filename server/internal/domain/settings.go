package domain

type ESuit string

const (
	Booby  ESuit = "♦"
	Chervy ESuit = "♥"
	Picky  ESuit = "♠"
	Tref   ESuit = "♣"
)

type ECard string

const (
	Seven ECard = "7"
	Eight ECard = "8"
	Nine  ECard = "9"
	Ten   ECard = "10"
	Jack  ECard = "J"
	Queen ECard = "Q"
	King  ECard = "K"
	Ace   ECard = "A"
)

type CardType struct {
	Type  ECard  `json:"type"`
	Name  string `json:"name"`
	Order byte   `json:"order"`
	Score byte   `json:"score"`
}
type CardSuit struct {
	Suit  ESuit  `json:"suit"`
	Name  string `json:"name"`
	Order byte   `json:"order"`
}

var CARD_TYPES = [8]CardType{
	{Type: Seven, Name: "7", Order: 1, Score: 0},
	{Type: Eight, Name: "8", Order: 2, Score: 0},
	{Type: Nine, Name: "9", Order: 3, Score: 0},
	{Type: Ten, Name: "10", Order: 6, Score: 10},
	{Type: Jack, Name: "J", Order: 8, Score: 2},
	{Type: Queen, Name: "Q", Order: 4, Score: 3},
	{Type: King, Name: "K", Order: 5, Score: 4},
	{Type: Ace, Name: "A", Order: 7, Score: 11},
}
var CARD_SUITS = [4]CardSuit{
	{Suit: Booby, Name: "♦", Order: 1},
	{Suit: Chervy, Name: "♥", Order: 2},
	{Suit: Picky, Name: "♠", Order: 3},
	{Suit: Tref, Name: "♣", Order: 4},
}
