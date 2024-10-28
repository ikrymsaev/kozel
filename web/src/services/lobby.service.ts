import { ILobby } from "../models/ILobby"
import { useAuthStore } from "../stores/auth.store"
import { useChatStore } from "../stores/chat.store"
import { useLobbyStore } from "../stores/lobby.store"
import { WS } from "./ws"

class LobbyService extends WS {

  
  /** Create a new lobby */
  public newLobby = async () => {
    const user = useAuthStore.getState().user
    if (!user) return

    const { id, name } = user;
    let url = "http://localhost:8080/lobby/new"
    url += "?user_id=" + id;
    url += "&username=" + name;
    try {
      const res = await fetch(url)
      const lobbyId = await res.json()
      return lobbyId
    } catch (e) {
      console.error('Failed to create lobby: ', e)
    }
  }

  /** Join a lobby  */
  public joinLobby = async (lobbyId: string) => {
    const user = useAuthStore.getState().user
    if (!user) return
    let url = "ws://localhost:8080/lobby/join/" + lobbyId
    url += "?user_id=" + user.id
    url += "&username=" + user.name
    try {
      this.conn(url)
      this.withConn((conn) => (conn.onmessage = this.onMessage))
      return 
    } catch (e) {
      console.error('Failed to join lobby: ', e)
    }
  }

  public leaveLobby = () => {
    this.disconn()
  }

  private onMessage = (e: { data: string }) => {
    const chatStore = useChatStore.getState()
    try {
      const message = JSON.parse(e.data) as string
      console.log(message)
      chatStore.addMessage(message)
    } catch (e) {
      console.error('Failed to parse message: ', e)
    }
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

    return () => es?.close()
  }
}

export const lobbyService = new LobbyService()