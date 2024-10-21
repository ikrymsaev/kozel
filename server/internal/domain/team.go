package domain

type Team struct {
	id   int
	A, B Player
}

func NewTeam(id int) Team {
	return Team{
		id: id,
	}
}

func (t *Team) AddPlayers(A Player, B Player) {
	t.A = A
	t.B = B

	t.A.Team = t
	t.B.Team = t
}
