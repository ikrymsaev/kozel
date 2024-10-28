package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Slot struct {
	order  byte
	player *User
}

type Lobby struct {
	Id      string
	Name    string
	OwnerId string
	Players map[string]*User
	slots   [4]Slot
}

func NewLobby(ownerId string, name string) *Lobby {
	return &Lobby{
		Id:      uuid.New().String(),
		Name:    name,
		OwnerId: ownerId,
		Players: make(map[string]*User),
		slots:   [4]Slot{{order: 1}, {order: 2}, {order: 3}, {order: 4}},
	}
}

// Подключение игрока к лобби
func (l *Lobby) ConnectPlayer(user *User) error {
	// Поиск свободного слота
	var freeSlot *Slot
	for _, slot := range l.slots {
		if slot.player == nil {
			freeSlot = &slot
			break
		}
	}
	if freeSlot == nil {
		return fmt.Errorf("lobby has no free slots")
	}
	freeSlot.player = user

	// TODO: Ws notify

	return nil
}
