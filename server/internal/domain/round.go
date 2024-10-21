package domain

import "fmt"

type Round struct {
	isFirstRound    bool       // Первый раунд
	firstStepPlayer *Player    // Игрок который ходит первым
	praiser         *Player    // Игрок который хвалит
	deck            Deck       // Колода
	trump           *ESuit     // Козырь
	bribes          [2][]*Card // Кортеж взяток (комада А, команда В)
	stakes          []*Stake   // Коны
	game            *Game      // Игра
}

// Результат раунда
type RoundResult struct {
	winner *Team // Команда победитель
	scrore byte  // Заработанное количество очков
}

func NewRound(game *Game) Round {
	return Round{
		isFirstRound: len(game.rounds) == 0,
		deck:         NewDeck(),
		game:         game,
		bribes:       [2][]*Card{},
		stakes:       []*Stake{},
	}
}

func (r *Round) Play() RoundResult {
	/* Подготовка колоды */
	fmt.Printf("Deck: %v\n\n", r.deck.CardsString())
	fmt.Printf("Shuffling deck...\n")
	r.deck.Shuffle()
	fmt.Printf("Deck: %v\n\n", r.deck.CardsString())
	r.dealCards()

	/* Хвалим козырь */
	println("==================================")
	fmt.Printf("Round %d\n", 1)
	fmt.Printf("Praiser: %s\n", r.praiser.Name)
	trump := r.praiser.PraiseTrump()
	r.trump = trump
	r.deck.SetTrump(trump)
	fmt.Printf("Trump: %s\n", *trump)
	/* Определение первого хода */
	r.firstStepPlayer = r.getFirstStepPlayer()
	fmt.Printf("First step player: %s\n", r.firstStepPlayer.Name)
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
	stake := NewStake(r, r.firstStepPlayer)
	r.stakes = append(r.stakes, &stake)
	stakeResult := stake.Start()
	fmt.Printf("Stake result >>> %s\n", stakeResult.winner.Name)
	fmt.Printf("Stake bribes: %v\n\n\n", stakeResult.bribe)

	if stakeResult.winner == &r.game.teams[0].A || stakeResult.winner == &r.game.teams[0].B {
		r.bribes[0] = append(r.bribes[0], stakeResult.bribe...)
	} else {
		r.bribes[1] = append(r.bribes[1], stakeResult.bribe...)
	}
	r.firstStepPlayer = stakeResult.winner
}

func (r *Round) getWinner() (*Team, byte) {
	var score_a byte = GetCardsScore(r.bribes[0])
	var score_b byte = GetCardsScore(r.bribes[1])

	fmt.Printf("Team A bribes score: %d\n", score_a)
	fmt.Printf("Team B bribes score: %d\n", score_b)

	if score_a > score_b {
		isPraiser := r.praiser.Team == &r.game.teams[0]
		return &r.game.teams[0], GetWinnerScores(isPraiser, score_a)
	}
	if score_a < score_b {
		isPraiser := r.praiser.Team == &r.game.teams[1]
		return &r.game.teams[1], GetWinnerScores(isPraiser, score_b)
	}
	return nil, 0
}

func (r *Round) getFirstStepPlayer() *Player {
	if r.isFirstRound {
		return &r.game.teams[0].A
	}
	return &r.game.teams[1].A
}

func (r *Round) dealCards() {
	cards := &r.deck.Cards
	for i := range cards {
		card := &cards[i]
		isPraiserCard := (card.CardSuit.Suit == Tref && card.CardType.Type == Jack) // TODO: refactor

		if i < 8 {
			r.game.teams[0].A.GetCard(card)
			if isPraiserCard {
				r.praiser = &r.game.teams[0].A
			}
		} else if i < 16 {
			r.game.teams[0].B.GetCard(card)
			if isPraiserCard {
				r.praiser = &r.game.teams[0].B
			}
		} else if i < 24 {
			r.game.teams[1].A.GetCard(card)
			if isPraiserCard {
				r.praiser = &r.game.teams[1].A
			}
		} else {
			r.game.teams[1].B.GetCard(card)
			if isPraiserCard {
				r.praiser = &r.game.teams[1].B
			}
		}
	}

	fmt.Printf("Dealed to %s\n", r.game.teams[0].A.HandString())
	fmt.Printf("Dealed to %s\n", r.game.teams[0].B.HandString())
	fmt.Printf("Dealed to %s\n", r.game.teams[1].A.HandString())
	fmt.Printf("Dealed to %s\n", r.game.teams[1].B.HandString())
}
