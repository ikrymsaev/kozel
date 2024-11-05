package domain

import (
	"fmt"
)

type Player struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Hand     []*Card `json:"hand"`
	Position byte    `json:"position"`
	Team     *Team   `json:"team"`
	User     *User   `json:"user"`
}

func NewPlayer(id string, name string, position byte, user *User) Player {
	return Player{
		Id:       id,
		Name:     name,
		Hand:     []*Card{},
		Position: position,
		User:     user,
	}
}

func (p *Player) IsBot() bool {
	return p.User == nil
}

func (p *Player) GetCard(card *Card) {
	p.Hand = append(p.Hand, card)
	card.SetOwner(p)
}

// Ход игрока
func (p *Player) Action(stake *Stake) *Card {
	var actionCard *Card
	if IsFirstStep(&stake.Table) {
		actionCard = p.firstStepLogic(stake) // Первый ход
	} else {
		actionCard = p.otherStepsLogic(stake) // Не первый ход
	}

	fmt.Printf("Bot: %v \n", p.Name)
	fmt.Printf("Hand %v \n", p.HandString())
	fmt.Printf("actionCard >>> %v \n\n", actionCard)

	p.removeCardFromHand(actionCard)
	return actionCard
}

func (p *Player) PlayerAction(cardId string) (*Card, error) {
	fmt.Printf("Hand: %v\n", p.Hand)
	targetCard := GetCardById(&p.Hand, cardId)
	fmt.Printf("Player %s: targetCard: %v\n", p.Name, targetCard)

	if targetCard == nil {
		return targetCard, fmt.Errorf("card no founded in hand")
	}
	p.removeCardFromHand(targetCard)

	return targetCard, nil
}

func (p *Player) removeCardFromHand(card *Card) {
	for i := 0; i < len(p.Hand); i++ {
		if p.Hand[i].Id == card.Id {
			copy(p.Hand[i:], p.Hand[i+1:])
			p.Hand = p.Hand[:len(p.Hand)-1]
			break
		}
	}
}

// Логика остальных ходов
func (p *Player) otherStepsLogic(stake *Stake) *Card {
	stackSuit := stake.GetStakeSuit()
	isTrumpStack := stackSuit == stake.round.Trump
	myTrumpCards := GetCardsBySuit(stake.round.Trump, &p.Hand)

	// Заход не по козырю
	if !isTrumpStack {
		suitsCards := GetCardsBySuit(stackSuit, &p.Hand)
		if len(suitsCards) > 0 {
			return suitsCards[0]
		}
		if len(myTrumpCards) > 0 {
			return myTrumpCards[0]
		}
		jacks := GetJacks(&p.Hand)
		if len(jacks) > 0 {
			return jacks[0]
		}
		return p.Hand[0]
	}

	/*
	 ? Кон по козырю
	*/
	winCard := GetWinCard(&stake.Table)
	isOurBribe := p.isOurBribe(&stake.Table)
	// Если взятка не наша
	if !isOurBribe {
		fmt.Println("!isOurBribe")
		// TODO Если смогу перебить
		if len(myTrumpCards) == 0 {
			return p.Hand[0]
		}
		isStartsWithJack := stake.Table[0].CardType.Type == Jack
		if isStartsWithJack {
			myOlderJack := GetOlderJack(&p.Hand)
			if myOlderJack != nil && myOlderJack.CardSuit.Order > winCard.CardSuit.Order {
				return myOlderJack
			} else {
				mySmallestTrump := GetSmallestScoreTrump(&p.Hand)
				if mySmallestTrump != nil {
					return mySmallestTrump
				}
				return GetSmallestScoreCard(&p.Hand)
			}
		}
		myOlderTrump := GetOlderTrump(&p.Hand)
		if myOlderTrump != nil && myOlderTrump.CardType.Order > winCard.CardType.Order {
			return myOlderTrump
		}
		myOlderJack := GetOlderJack(&p.Hand)
		if myOlderJack != nil {
			return myOlderJack
		}

		if myOlderTrump == nil {
			return GetSmallestScoreCard(&p.Hand)
		}
		return myOlderTrump
	}
	// Если взятка наша
	hasJacksInGame := len(p.getJacksInGame(stake)) > 0

	if hasJacksInGame {
		bestScoreSuitCard := GetBestScoreSuitCard(&p.Hand, stackSuit)
		if bestScoreSuitCard != nil {
			return bestScoreSuitCard
		}
	}
	myAce := p.FindCardInHand(Ace, *stackSuit)
	if myAce != nil {
		return myAce
	}
	myTen := p.FindCardInHand(Ten, *stackSuit)
	if myTen != nil {
		return myTen
	}

	return GetBestScoreCard(&p.Hand)
}

func (p *Player) isOurBribe(cards *[]*Card) bool {
	winCard := GetWinCard(cards)
	isOurBribe := winCard.Owner.Team == p.Team
	return isOurBribe
}

func (p *Player) FindCardInHand(cardType ECard, suit ESuit) *Card {
	for _, card := range p.Hand {
		if card.CardType.Type == cardType && card.CardSuit.Suit == suit {
			return card
		}
	}
	return nil
}

func (p *Player) getJacksInGame(stake *Stake) []*Card {
	jacks := make([]*Card, 0)

	for _, card := range stake.round.Deck.Cards {
		if card.CardType.Type == Jack && !card.IsUsed && card.Owner != p {
			jacks = append(jacks, &card)
		}
	}
	return jacks
}

// Логика первого хода
func (p *Player) firstStepLogic(stake *Stake) *Card {
	wePraiser := IsWePraiser(p, stake)
	/*
	 ? Если МЫ хвалили
	*/
	if wePraiser {
		fmt.Printf("We praise!\n")
		// Оцениваем ход с вальта
		stepWithJack, jackCard := p.jackLogic(stake, wePraiser)
		fmt.Printf("stepWithJack: %v %v\n", stepWithJack, jackCard)
		if stepWithJack {
			return jackCard
		}
		// Оцениваем ход с козыря
		stepWithTrump, trumpCard := p.trumpLogic(stake)
		fmt.Printf("stepWithTrump: %v %v\n", stepWithTrump, trumpCard)
		if stepWithTrump {
			return trumpCard
		}
	}
	/*
	 ? Если НЕ хвалили
	*/
	// Oцениваем ход с туза
	stepWithAce, aceCard := p.aceLogic(stake)
	if stepWithAce {
		return aceCard
	}
	// Оцениваем ход с 10
	stepWithTen, tenCard := p.tenLogic(stake)
	if stepWithTen {
		return tenCard
	}

	// Пытаемся сбросить карту
	return p.noVariantsLogic()
}

// Логика если есть 10
func (p *Player) tenLogic(stake *Stake) (bool, *Card) {
	handTens := GetTens(&p.Hand)
	if len(handTens) == 0 {
		return false, nil
	}

	// Считаем количество невышедших козырей
	totalTrumpsInGame := p.getTrumpsInGameCount(stake)
	if totalTrumpsInGame > 2 {
		return false, nil
	}

	for _, ten := range handTens {
		if ten.IsTrump {
			continue
		}
		// Если в игре есть туз то не рискуем 10
		isAceInDeck := stake.round.Deck.IsHasAce(&ten.CardSuit.Suit)
		isMyAce := IsHasAce(&ten.CardSuit.Suit, &p.Hand)
		if isAceInDeck && !isMyAce {
			continue
		}

		// Если в игре достаточно мастей, то рискуя ходим 10
		suitsInGame := len(stake.round.Deck.GetSuitsInGame(&ten.CardSuit.Suit)) - 1
		if suitsInGame > 2 {
			return true, ten
		}
	}

	return false, nil
}

// Логика если есть туз
func (p *Player) aceLogic(stake *Stake) (bool, *Card) {
	handAces := GetAces(&p.Hand)
	if len(handAces) == 0 {
		return false, nil
	}

	// Считаем количество невышедших козырей
	totalTrumpsInGame := p.getTrumpsInGameCount(stake)

	// Если в игре много козырей, то не рискуем
	if totalTrumpsInGame > 2 {
		return false, nil
	}

	for _, ace := range handAces {
		if ace.IsTrump {
			continue
		}
		// Если в игре достаточно мастей, то рискуя ходим тузом
		suitsInGame := len(stake.round.Deck.GetSuitsInGame(&ace.CardSuit.Suit)) - 1
		if suitsInGame > 1 {
			return true, ace
		}
	}

	return false, nil
}

// Логика если есть козыри в руке
func (p *Player) trumpLogic(stake *Stake) (bool, *Card) {
	trumpCards := GetTrumpCards(&p.Hand)
	if len(trumpCards) == 0 {
		return false, nil
	}
	myOlderTrump := GetOlderTrump(&p.Hand)

	// Если старший козырь бесполезен - ходим с него
	if myOlderTrump.CardType.Score == 0 {
		return true, myOlderTrump
	}

	hasJackInGame := stake.round.Deck.HasJackInGame()
	olderTrumpInGame := stake.round.Deck.GetOlderTrumpInGame()

	isOlderTrumpIsMine := myOlderTrump != nil && myOlderTrump.CardType.Type == olderTrumpInGame.CardType.Type

	// Если козырь - старший в игре
	if !hasJackInGame && isOlderTrumpIsMine {
		return true, myOlderTrump
	}

	// Если 10 старшая в руке, но не в игре
	if myOlderTrump.CardType.Type == Ten {
		// Если есть козырь кроме 10, то берем его
		smallestScoreTrump := GetSmallestScoreTrump(&p.Hand)
		if (smallestScoreTrump != nil) && (smallestScoreTrump.CardType.Score == 0) {
			return true, smallestScoreTrump
		}
	}

	return false, nil
}

// Логика если есть валет
func (p *Player) jackLogic(stake *Stake, wePraiser bool) (bool, *Card) {
	myOlderJack := GetOlderJack(&p.Hand) // Старший валет в руке
	if myOlderJack == nil {
		return false, nil
	}
	stepWithJack := IsMyJackIsOlder(myOlderJack, &stake.round.Deck, wePraiser) // Проверяем страший ли в игре
	return stepWithJack, myOlderJack
}

// Логика остальных ходов
func (p *Player) noVariantsLogic() *Card {
	// Сначала есть ли что-то кроме тузов тыщи и вальта
	uselessCards := GetUselessCards(&p.Hand)
	if len(uselessCards) > 0 {
		return GetSmallestScoreCard(&uselessCards)
	}

	// Есть ли что-то кроме козырей
	noneTrumpCards := GetNoneTrumpCards(&p.Hand)
	if len(noneTrumpCards) > 0 {
		return GetSmallestScoreCard(&noneTrumpCards)
	}

	// Если нет ни одного хорошего хода, то берем случайную карту
	return GetRandomCard(&p.Hand)
}

// Захвалить масть
func (p *Player) PraiseTrump() *ESuit {
	choosedSuit := &p.Hand[0].CardSuit.Suit

	// Считаем количество карт одной масти
	suitsCount := make(map[ESuit]int)
	// Максимальное количество карт одной масти которое можно захвалить
	maxCountCards := 0
	for _, card := range p.Hand {
		if card.CardType.Type != Jack {
			suitsCount[card.CardSuit.Suit]++
			if suitsCount[card.CardSuit.Suit] > maxCountCards {
				maxCountCards = suitsCount[card.CardSuit.Suit]
			}
		}
	}

	// Смотрим масти с максимальным количеством карт
	suitsForChoose := []*ESuit{}
	for suit, count := range suitsCount {
		if count == maxCountCards {
			suitsForChoose = append(suitsForChoose, &suit)
		}
	}

	//? Если есть одна масть с максимальным количеством карт, то хвалим ее
	if len(suitsForChoose) == 1 {
		return suitsForChoose[0]
	}

	totalScore := make(map[ESuit]byte)

	// Считаем сумму очков для каждой масти
	for _, card := range p.Hand {
		if card.CardType.Type != Jack {
			totalScore[card.CardSuit.Suit] += card.CardType.Score
		}
	}

	var maxScore byte = 0
	// Смотрим масть с максимальным суммой очков
	for suit, score := range totalScore {
		if score > maxScore {
			maxScore = score
			choosedSuit = &suit
		}
	}

	return choosedSuit
}

func (p *Player) HandString() string {
	return fmt.Sprintf("Player %d: %s", p.Id, p.Hand)
}

// Считаем количество невышедших козырей
func (p *Player) getTrumpsInGameCount(stake *Stake) int {
	trumpsInGameCount := len(stake.round.Deck.GetTrumpsInGame())
	jacksInGameCount := len(stake.round.Deck.GetJacksInGame())
	myTrumpsCount := len(GetTrumpCards(&p.Hand))
	myJacksCount := len(GetJacks(&p.Hand))
	totalJacksInGame := jacksInGameCount - myJacksCount

	totalTrumpsInGame := (trumpsInGameCount - myTrumpsCount) + totalJacksInGame

	return totalTrumpsInGame
}
