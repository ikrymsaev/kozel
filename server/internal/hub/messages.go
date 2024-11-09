package hub

// @Messages send to the client
type EHubMessage int

const (
	EHubMessageNewLobby    EHubMessage = iota
	EHubMessageRemoveLobby EHubMessage = iota
)

type LobbyMessage struct {
	Type EHubMessage `json:"type"`
	Id   string      `json:"id"`
}
