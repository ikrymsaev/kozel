import { useLobbyStore } from "../stores/lobby.store"

const MOCK_USER_ID = "dag87a5dg"
const MOCK_USERNAME = 'hellfever'

class LobbyService {

  public newGame = async () => {
    return await fetch('http://localhost:8080/ws/new_game', {
      method: 'POST',
      body: JSON.stringify({ id: MOCK_USER_ID, name: MOCK_USERNAME }),
    })
  }

  public getGames = async () => {
    const res = await fetch('http://localhost:8080/ws/get_games')
    const games = await res.json()
    useLobbyStore.setState({ activeGames: games })
  }
}

export const lobbyService = new LobbyService()