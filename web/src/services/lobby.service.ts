import { useAuthStore } from "../stores/auth.store"
import { useLobbyStore } from "../stores/lobby.store"

class LobbyService {

  public newGame = async () => {
    const user = useAuthStore.getState().user
    if (!user) return

    const { id, name } = user;
    return await fetch('http://localhost:8080/ws/new_game', {
      method: 'POST',
      body: JSON.stringify({ id, name }),
    })
  }

  public getGames = async () => {
    const res = await fetch('http://localhost:8080/ws/get_games')
    const games = await res.json()
    useLobbyStore.setState({ activeGames: games })
  }
}

export const lobbyService = new LobbyService()