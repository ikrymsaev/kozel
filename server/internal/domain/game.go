package domain

import "fmt"

type Game struct {
	teams  [2]Team
	rounds []Round
	score  [2]byte
}

func NewGame() Game {
	teamA, teamB := createTeams()
	return Game{
		rounds: make([]Round, 0),
		score:  [2]byte{0, 0},
		teams:  [2]Team{teamA, teamB},
	}
}

func (g *Game) Start() {
	println("Starting game...")

	round := NewRound(g)
	g.rounds = append(g.rounds, round)

	result := round.Play()

	fmt.Printf("Round result >>> %s\n", result.winner.A.Name)
}

func createTeams() (Team, Team) {
	teamA := NewTeam(1)
	teamA.AddPlayers(NewPlayer(1, "Player 1", 1), NewPlayer(3, "Player 3", 3))
	teamB := NewTeam(2)
	teamB.AddPlayers(NewPlayer(2, "Player 2", 2), NewPlayer(4, "Player 4", 4))

	return teamA, teamB
}
