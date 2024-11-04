package domain

import (
	"fmt"
	"go-kozel/pkg/utils"
)

type Game struct {
	Teams        [2]Team
	Rounds       []Round
	CurrentRound Round
	Score        [2]byte
}

func NewGame(lobby *Lobby) Game {
	teamA, teamB := createTeams(lobby.Slots)
	return Game{
		Rounds: make([]Round, 0),
		Score:  [2]byte{0, 0},
		Teams:  [2]Team{teamA, teamB},
	}
}

func (g *Game) Start() {
	println("Starting game...")

	round := NewRound(g)
	g.CurrentRound = round
	g.CurrentRound.Init()
	g.Rounds = append(g.Rounds, round)
	fmt.Printf("Game FirstStepPlayer %v\n", round.FirstStepPlayer)
}

func (g *Game) GetPlayers() [4]Player {
	var players [4]Player
	players[0] = g.Teams[0].A
	players[1] = g.Teams[1].A
	players[2] = g.Teams[0].B
	players[3] = g.Teams[1].B
	return players
}

func createTeams(slots [4]Slot) (Team, Team) {
	p1 := NewPlayer("1", utils.GetRandomName(), 1, nil)
	p2 := NewPlayer("2", utils.GetRandomName(), 2, nil)
	p3 := NewPlayer("3", utils.GetRandomName(), 3, nil)
	p4 := NewPlayer("4", utils.GetRandomName(), 4, nil)

	slot1p := slots[0].Player
	if slot1p != nil {
		p1 = NewPlayer(slot1p.ID, slot1p.Username, 1, slot1p)
	}

	slot2p := slots[1].Player
	if slot2p != nil {
		p2 = NewPlayer(slot2p.ID, slot2p.Username, 2, slot2p)
	}

	slot3p := slots[2].Player
	if slot3p != nil {
		p3 = NewPlayer(slot3p.ID, slot3p.Username, 3, slot3p)
	}

	slot4p := slots[3].Player
	if slot4p != nil {
		p4 = NewPlayer(slot4p.ID, slot4p.Username, 4, slot4p)
	}

	teamA := NewTeam(1)
	teamA.AddPlayers(p1, p3)
	teamB := NewTeam(2)
	teamB.AddPlayers(p2, p4)

	return teamA, teamB
}
