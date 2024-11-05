package domain

type Team struct {
	Id   byte
	A, B Player
}

func NewTeam(id byte) Team {
	return Team{
		Id: id,
	}
}

func (t *Team) AddPlayers(A Player, B Player) {
	t.A = A
	t.B = B

	t.A.Team = t
	t.B.Team = t
}
