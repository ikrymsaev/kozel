import { toast } from "react-toastify"
import { useAuthStore } from "../stores/auth.store"
import { useChatStore } from "../stores/chat.store"
import { useLobbyStore } from "../stores/lobby.store"
import { EWSAction } from "../api/ws/ws.actions"
import { TWsService, wsService } from "../api/ws/ws.service"
import { EWSMessage, TConnectionMsg, TErrorMsg, TUpdateSlotsMsg } from "../api/ws/ws.messages"
import { getApiUrl } from "@/shared/utils/get-api-url"
import axios from "axios"
import { ILobby } from "@/models/ILobby"

class LobbyService {
  private readonly ws: TWsService
  constructor() {
    this.ws = wsService
    this.ws.listen(EWSMessage.Connection, this.onConnect)
    this.ws.listen(EWSMessage.UpdateSlots, this.onUpdateSlots)
    this.ws.listen(EWSMessage.Error, this.onError)
  }

  public fetchLobbyList = async () => {
    const url = getApiUrl() + "/api/lobby/list"
    const res = await fetch(url)
    const games = await res.json()
    useLobbyStore.setState({ activeGames: games })
  }

  public fetchNewLobby = async (id: string) => {
    const url = getApiUrl() + "/api/lobby/" + id
    const { data: lobby } = await axios.get<ILobby>(url)
    useLobbyStore.getState().addLobby(lobby)
  }

  /** Create a new lobby */
  public newLobby = async () => {
    const user = useAuthStore.getState().user
    if (!user) return

    const url = getApiUrl() + "/api/lobby/new"
    try {
      const { data: lobbyId } = await axios.get<{ data: string }>(
        url,
        {headers: {'Authorization': `Bearer ${useAuthStore.getState().token}`}}
      )
      return lobbyId
    } catch (e) {
      console.error('Failed to create lobby: ', e)
    }
  }

  /** Join a lobby  */
  public joinLobby = async (lobbyId: string) => {
    const user = useAuthStore.getState().user
    const token = useAuthStore.getState().token
    if (!user || !token) return
    
    const url = getApiUrl() + "/api/lobby/join/" + lobbyId + "/" + token
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