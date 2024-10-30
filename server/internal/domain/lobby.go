package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Slot struct {
	Order  byte  `json:"order"`
	Player *User `json:"player"`
}

type Lobby struct {
	Id      string
	Name    string
	OwnerId string
	Slots   [4]Slot
}

func NewLobby(ownerId string, name string) *Lobby {
	return &Lobby{
		Id:      uuid.New().String(),
		Name:    name,
		OwnerId: ownerId,
		Slots:   [4]Slot{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}},
	}
}

func (l *Lobby) GetSlots() [4]Slot {
	return l.Slots
}

// Подключение игрока к лобби
func (l *Lobby) ConnectPlayer(user *User) error {
	// Поиск свободного слота
	for index, slot := range l.Slots {
		if slot.Player == nil {
			l.Slots[index].Player = user
			return nil
		}
	}
	return fmt.Errorf("lobby has no free slots")
}

// Отключение игрока от лобби
func (l *Lobby) DisconnectPlayer(user *User) error {
	for index, slot := range l.Slots {
		if slot.Player.ID == user.ID {
			l.Slots[index].Player = nil
			return nil
		}
	}
	return fmt.Errorf("player not found")
}
