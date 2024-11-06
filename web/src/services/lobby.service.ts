import { toast } from "react-toastify"
import { ILobby } from "../models/ILobby"
import { useAuthStore } from "../stores/auth.store"
import { useChatStore } from "../stores/chat.store"
import { useLobbyStore } from "../stores/lobby.store"
import { EWSAction } from "../api/ws/ws.actions"
import { TWsService, wsService } from "../api/ws/ws.service"
import { EWSMessage, TConnectionMsg, TErrorMsg, TUpdateSlotsMsg } from "../api/ws/ws.messages"

class LobbyService {
  private readonly ws: TWsService
  constructor() {
    this.ws = wsService
    this.ws.listen(EWSMessage.Connection, this.onConnect)
    this.ws.listen(EWSMessage.UpdateSlots, this.onUpdateSlots)
    this.ws.listen(EWSMessage.Error, this.onError)
  }

  public getLobbyList = async () => {
    const res = await fetch('http://localhost:8080/hub/lobbies')
    const games = await res.json()
    useLobbyStore.setState({ activeGames: games })
  }

  public watchLobbies = () => {
    this.getLobbyList()

    const { addLobby  } = useLobbyStore.getState()

    const connect = () => {
      const es = new EventSource('http://localhost:8080/hub/watch_lobbies')
      console.log("SSE opened, watching for lobbies...")
      es.addEventListener('new_lobby', (event) => {
        const data = JSON.parse(event.data) as ILobby
        console.log("SSE new_lobby: ", data)
        addLobby(data)
      })
  
      es.addEventListener('remove_lobby', (event) => {
        const lobbyId = event.data
        console.log("SSE remove_lobby: ", lobbyId)
        useLobbyStore.getState().removeLobby(lobbyId)
      })
      return {
        es,
        close: () => {
          es.close()
        },
      }
    }
    const { es, close } = connect()
    es.onerror = () => {
      setTimeout(() => {
        close()
        connect()
      }, 3000)
    }

    return close
  }

  /** Create a new lobby */
  public newLobby = async () => {
    const user = useAuthStore.getState().user
    if (!user) return

    const { id, username } = user;
    let url = "http://localhost:8080/lobby/new"
    url += "?user_id=" + id;
    url += "&username=" + username;
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
    url += "&username=" + user.username
    try {
      return this.ws.connect(url)
    } catch (e) {
      console.error('Failed to join lobby: ', e)
    }
  }

  /** Leave a lobby */
  public leaveLobby = () => {
    this.ws.disconnect()
  }

  public moveSlot = (from: number, to: number) => {
    this.ws.send({ type: EWSAction.MoveSlot, from, to })
  }

  private onError = (msg: TErrorMsg) => {
    console.error(msg)
    toast(msg.error, { type: 'error' })
  }

  private onUpdateSlots = (m: TUpdateSlotsMsg) => {
    console.log({ m })
    const lobbyStore = useLobbyStore.getState()
    const { slots } = m
    lobbyStore.updateSlots(slots)
  }

  private onConnect = (msg: TConnectionMsg) => {
    const chatStore = useChatStore.getState()
    const me = useAuthStore.getState().user
    const { isConnected, user } = msg
    const isMe = user?.id === me?.id
    const username = !user ? "unknown: " : `${user.username}`
    const connMessage = `${username}${isConnected ? ' joined to lobby' : ' left from lobby'}`
    if (isMe && isConnected) {
      chatStore.reset()
    }
    if (!isMe) toast(connMessage)
  }
}

export const lobbyService = new LobbyService()