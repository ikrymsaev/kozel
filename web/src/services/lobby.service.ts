import { ILobby } from "../models/ILobby"
import { useAuthStore } from "../stores/auth.store"
import { useLobbyStore } from "../stores/lobby.store"

class LobbyService {

  public newGame = async () => {
    const user = useAuthStore.getState().user
    if (!user) return

    const { id, name } = user;
    return await fetch('http://localhost:8080/hub/new_lobby', {
      method: 'POST',
      body: JSON.stringify({ id, name }),
    })
  }

  public getLobbyList = async () => {
    const res = await fetch('http://localhost:8080/hub/lobbies')
    const games = await res.json()
    useLobbyStore.setState({ activeGames: games })
  }

  public watchLobbies = () => {
    this.getLobbyList()

    const { addLobby } = useLobbyStore.getState()
    const es = new EventSource('http://localhost:8080/hub/watch_lobbies')
    console.log("SSE opened, watching for lobbies...")
    es.addEventListener('new_lobby', (event) => {
      const data = JSON.parse(event.data) as ILobby
      console.log("SSE new_lobby: ", data)
      addLobby(data)
    })

    return es.close
  }
}

export const lobbyService = new LobbyService()