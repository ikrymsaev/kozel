package domain

import "fmt"

type Round struct {
	IsFirstRound    bool       `json:"isFirstRound"`
	FirstStepPlayer *Player    `json:"firstStepPlayer"`
	Praiser         *Player    `json:"praiser"`
	Deck            Deck       `json:"deck"`
	Trump           *ESuit     `json:"trump"`
	Bribes          [2][]*Card `json:"bribes"`
	Stakes          []*Stake   `json:"stake"`
	Game            *Game      `json:"game"`
	CurrentStake    *Stake     `json:"currentStake"`
}

// Результат раунда
type RoundResult struct {
	winner *Team // Команда победитель
	scrore byte  // Заработанное количество очков
}

func NewRound(game *Game) Round {
	return Round{
		IsFirstRound: len(game.Rounds) == 0,
		Deck:         NewDeck(),
		Game:         game,
		Bribes:       [2][]*Card{},
		Stakes:       []*Stake{},
	}
}

func (r *Round) Init() {
	r.Deck.Shuffle()
	r.dealCards()
	firstPlayer := r.getFirstStepPlayer()
	r.FirstStepPlayer = firstPlayer
	r.Game.Stage = StagePraising
	fmt.Printf("StagePraising: %v\n\n", r.FirstStepPlayer)
}

func (r *Round) SetTrump(trump *ESuit) {
	r.Trump = trump
	r.Deck.SetTrump(trump)
}

func (r *Round) InitStake() {
	stake := NewStake(r, r.FirstStepPlayer)
	r.CurrentStake = &stake
	r.Stakes = append(r.Stakes, r.CurrentStake)
}

func (r *Round) Play() RoundResult {
	/* Подготовка колоды */
	fmt.Printf("Deck: %v\n\n", r.Deck.CardsString())
	fmt.Printf("Shuffling deck...\n")
	r.Deck.Shuffle()
	fmt.Printf("Deck: %v\n\n", r.Deck.CardsString())
	r.dealCards()

	/* Хвалим козырь */
	println("==================================")
	fmt.Printf("Round %d\n", 1)
	fmt.Printf("Praiser: %s\n", r.Praiser.Name)
	trump := r.Praiser.PraiseTrump()
	r.Trump = trump
	r.Deck.SetTrump(trump)
	fmt.Printf("Trump: %s\n", *trump)
	/* Определение первого хода */
	r.FirstStepPlayer = r.getFirstStepPlayer()
	fmt.Printf("First step player: %s\n", r.FirstStepPlayer.Name)
	println("==================================")

	for i := 0; i < 8; i++ {
		r.initStake()
	}

	winner, roundScore := r.getWinner()
	println("==================================")
	if winner != nil {
		fmt.Printf("Winner Team: %d\n", winner.id)
	}
	fmt.Printf("Round score: %d\n", roundScore)
	println("==================================")

	return RoundResult{
		winner: winner,
		scrore: roundScore,
	}
}

func (r *Round) initStake() {
	/* Запуск нового кона */
	stake := NewStake(r, r.FirstStepPlayer)
	r.Stakes = append(r.Stakes, &stake)
	stakeResult := stake.Start()
	fmt.Printf("Stake result >>> %s\n", stakeResult.Winner.Name)
	fmt.Printf("Stake bribes: %v\n\n\n", stakeResult.Bribe)

	if stakeResult.Winner == &r.Game.Teams[0].A || stakeResult.Winner == &r.Game.Teams[0].B {
		r.Bribes[0] = append(r.Bribes[0], stakeResult.Bribe...)
	} else {
		r.Bribes[1] = append(r.Bribes[1], stakeResult.Bribe...)
	}
	r.FirstStepPlayer = stakeResult.Winner
}

func (r *Round) getWinner() (*Team, byte) {
	var score_a byte = GetCardsScore(r.Bribes[0])
	var score_b byte = GetCardsScore(r.Bribes[1])

	fmt.Printf("Team A bribes score: %d\n", score_a)
	fmt.Printf("Team B bribes score: %d\n", score_b)

	if score_a > score_b {
		isPraiser := r.Praiser.Team == &r.Game.Teams[0]
		return &r.Game.Teams[0], GetWinnerScores(isPraiser, score_a)
	}
	if score_a < score_b {
		isPraiser := r.Praiser.Team == &r.Game.Teams[1]
		return &r.Game.Teams[1], GetWinnerScores(isPraiser, score_b)
	}
	return nil, 0
}

func (r *Round) getFirstStepPlayer() *Player {
	if r.IsFirstRound {
		return &r.Game.Teams[0].A
	}
	return &r.Game.Teams[1].A
}

func (r *Round) dealCards() {
	cards := &r.Deck.Cards
	for i := range cards {
		card := &cards[i]
		isPraiserCard := (card.CardSuit.Suit == Tref && card.CardType.Type == Jack) // TODO: refactor

		if i < 8 {
			r.Game.Teams[0].A.GetCard(card)
			if isPraiserCard {
				r.Praiser = &r.Game.Teams[0].A
			}
		} else if i < 16 {
			r.Game.Teams[0].B.GetCard(card)
			if isPraiserCard {
				r.Praiser = &r.Game.Teams[0].B
			}
		} else if i < 24 {
			r.Game.Teams[1].A.GetCard(card)
			if isPraiserCard {
				r.Praiser = &r.Game.Teams[1].A
			}
		} else {
			r.Game.Teams[1].B.GetCard(card)
			if isPraiserCard {
				r.Praiser = &r.Game.Teams[1].B
			}
		}
	}

	fmt.Printf("Dealed to %s\n", r.Game.Teams[0].A.HandString())
	fmt.Printf("Dealed to %s\n", r.Game.Teams[0].B.HandString())
	fmt.Printf("Dealed to %s\n", r.Game.Teams[1].A.HandString())
	fmt.Printf("Dealed to %s\n", r.Game.Teams[1].B.HandString())
}
