package domain

import (
	"math/rand"
)

// Является ли первым ходом
func IsFirstStep(cards *[]*Card) bool {
	return len(*cards) == 0
}

func IsWePraiser(player *Player, stake *Stake) bool {
	praiserTeam := stake.round.praiser.Team
	return player.Team == praiserTeam
}

func GetTrumpCards(cards *[]*Card) []*Card {
	var trumpCards []*Card

	for _, card := range *cards {
		if card.IsTrump && card.CardType.Type != Jack {
			trumpCards = append(trumpCards, card)
		}
	}

	return trumpCards
}

func GetJacks(cards *[]*Card) []*Card {
	var jacks []*Card

	for _, card := range *cards {
		if card.CardType.Type == Jack {
			jacks = append(jacks, card)
		}
	}
	return jacks
}

func GetOlderTrump(cards *[]*Card) *Card {
	var olderTrump *Card

	for _, card := range *cards {
		if card.IsTrump && card.CardType.Type != Jack {
			isOlder := olderTrump == nil || card.CardType.Order > olderTrump.CardType.Order
			if isOlder {
				olderTrump = card
			}
		}
	}
	return olderTrump
}

func GetSmallestScoreTrump(cards *[]*Card) *Card {
	var smallestTrump *Card

	for _, card := range *cards {
		if card.IsTrump && card.CardType.Type != Jack {
			isSmallest := smallestTrump == nil || card.CardType.Score < smallestTrump.CardType.Score
			if isSmallest {
				smallestTrump = card
			}
		}
	}
	return smallestTrump
}

func GetOlderJack(cards *[]*Card) *Card {
	var olderJack *Card

	for _, card := range *cards {
		if card.CardType.Type != Jack {
			continue
		}
		isOlder := olderJack == nil || card.CardSuit.Order > olderJack.CardSuit.Order
		if isOlder {
			olderJack = card
		}
	}
	return olderJack
}

func IsMyJackIsOlder(myJack *Card, deck *Deck, wePraiser bool) bool {
	jacksInGame := deck.GetJacksInGame()
	var olderJackInBribes *Card

	for _, card := range jacksInGame {
		if card.CardType.Type == Jack && (wePraiser && card.CardSuit.Suit != Tref) { // Если мы хвалили то не считаем старшего вальта
			isOlder := olderJackInBribes == nil || card.CardSuit.Order >= olderJackInBribes.CardSuit.Order
			if isOlder {
				olderJackInBribes = card
			}
		}
	}
	if olderJackInBribes == nil {
		return true
	}

	return myJack.CardSuit.Order >= olderJackInBribes.CardSuit.Order
}

func GetAces(cards *[]*Card) []*Card {
	var aces []*Card
	for _, card := range *cards {
		if card.CardType.Type == Ace {
			aces = append(aces, card)
		}
	}
	return aces
}

func GetTens(cards *[]*Card) []*Card {
	var tens []*Card
	for _, card := range *cards {
		if card.CardType.Type == Ten {
			tens = append(tens, card)
		}
	}
	return tens
}

func IsHasAce(suit *ESuit, cards *[]*Card) bool {
	for _, card := range *cards {
		if card.CardType.Type == Ace && card.CardSuit.Suit == *suit {
			return true
		}
	}
	return false
}

func GetUselessCards(cards *[]*Card) []*Card {
	var uselessCards []*Card
	for _, card := range *cards {
		if !card.IsTrump && card.CardType.Type != Jack && card.CardType.Type != Ace && card.CardType.Type != Ten {
			uselessCards = append(uselessCards, card)
		}
	}
	return uselessCards
}

func GetNoneTrumpCards(cards *[]*Card) []*Card {
	var noneTrumpCards []*Card
	for _, card := range *cards {
		if !card.IsTrump {
			noneTrumpCards = append(noneTrumpCards, card)
		}
	}
	return noneTrumpCards
}

func GetSmallestScoreCard(cards *[]*Card) *Card {
	var smallestCard *Card
	for _, card := range *cards {
		if card.CardType.Type != Jack {
			isSmallest := smallestCard == nil || card.CardType.Score < smallestCard.CardType.Score
			if isSmallest {
				smallestCard = card
			}
		}
	}
	return smallestCard
}

func GetRandomCard(cards *[]*Card) *Card {
	rand.Shuffle(len(*cards), func(i, j int) {
		(*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
	})
	return (*cards)[0]
}

func GetWinCard(cards *[]*Card) *Card {
	if len(*cards) == 0 {
		return nil
	}
	winCard := (*cards)[0]

	for _, card := range *cards {
		winCard = GetOlderCard(winCard, card)
	}

	return winCard
}

func GetOlderCard(card_1 *Card, card_2 *Card) *Card {
	if card_1.IsTrump {
		if !card_2.IsTrump {
			return card_1
		}
		if card_1.CardType.Type == Jack {
			if card_2.CardType.Type != Jack {
				return card_1
			}
			if card_2.CardSuit.Order > card_1.CardSuit.Order {
				return card_2
			}
			return card_1
		}
		if card_2.CardType.Type == Jack {
			return card_2
		}
		if card_2.CardType.Order > card_1.CardType.Order {
			return card_2
		}
		return card_1
	}
	if card_2.IsTrump {
		return card_2
	}
	if card_2.CardSuit.Suit != card_1.CardSuit.Suit {
		return card_1
	}

	if card_2.CardType.Order > card_1.CardType.Order {
		return card_2
	}
	return card_1
}

func GetCardsBySuit(suit *ESuit, cards *[]*Card) []*Card {
	var cardsBySuit []*Card
	for _, card := range *cards {
		if card.CardSuit.Suit == *suit && card.CardType.Type != Jack {
			cardsBySuit = append(cardsBySuit, card)
		}
	}
	return cardsBySuit
}

func GetBestScoreSuitCard(cards *[]*Card, suit *ESuit) *Card {
	var bestScoreCard *Card
	for _, card := range *cards {
		isJack := card.CardType.Type == Jack
		isAce := card.CardType.Type == Ace
		isTen := card.CardType.Type == Ten
		isSameSuit := card.CardSuit.Suit == *suit && !isJack
		if isJack || isAce || isTen || !isSameSuit {
			continue
		}
		if bestScoreCard == nil || card.CardType.Score > bestScoreCard.CardType.Score {
			bestScoreCard = card
		}
	}
	return bestScoreCard
}

func GetBestScoreCard(cards *[]*Card) *Card {
	var bestScoreCard *Card
	for _, card := range *cards {
		if bestScoreCard == nil || card.CardType.Score > bestScoreCard.CardType.Score {
			bestScoreCard = card
		}
	}
	return bestScoreCard
}

func GetCardsScore(cards []*Card) byte {
	var scores byte = 0
	for _, card := range cards {
		scores += card.CardType.Score
	}
	return scores
}

// TODO Сделать все сравнения карт через одну функцию
// TODO func GetOlderCards
