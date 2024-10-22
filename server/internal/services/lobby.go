package services

import (
	"go-kozel/internal/domain"

	"github.com/google/uuid"
)

type Lobby struct {
	Id      string
	Name    string
	OwnerId string
	Players map[string]*domain.User
}

func NewLobby(ownerId string, name string) *Lobby {
	return &Lobby{
		Id:      uuid.New().String(),
		Name:    name,
		OwnerId: ownerId,
		Players: make(map[string]*domain.User),
	}
}
