import { toast } from "react-toastify"
import { ILobby, ILobbySlot } from "../models/ILobby"
import { useAuthStore } from "../stores/auth.store"
import { useChatStore } from "../stores/chat.store"
import { useLobbyStore } from "../stores/lobby.store"
import { WS } from "./ws"

class LobbyService extends WS {

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
      this.conn(url)
      this.withConn((conn) => (conn.onmessage = this.onMessage))
      return 
    } catch (e) {
      console.error('Failed to join lobby: ', e)
    }
  }

  /** Leave a lobby */
  public leaveLobby = () => {
    this.disconn()
  }

  public moveSlot = (from: number, to: number) => {
    this.withConn((conn) => conn.send(
      JSON.stringify({
        type: EMsgType.MoveSlot,
        from,
        to
      })
    ))
  }

  public startGame = () => {
    console.log('start game')
  }

  public sendChatMessage = (message: string) => {
    const user = useAuthStore.getState().user
    if (!user) return
    this.withConn((conn) => conn.send(
      JSON.stringify({
        type: EMsgType.Chat,
        data: { message }
      })
    ))
  }

  private onMessage = (e: { data: string }) => {
    try {
      const m = JSON.parse(e.data) as TWsMessage
      console.log("onMessage: ", m)
      if (isConnectionMsg(m)) this.onConnectMessage(m)
      else if (isChatMsg(m)) this.onChatMessage(m)
      else if (isUpdateMsg(m)) this.onUpdateMessage(m)
      else console.error('Unknown message: ', m)
    } catch (e) {
      console.error('Failed to parse message: ', e)
    }
  }

  private onUpdateMessage = (m: TUpdateMsg) => {
    console.log({ m })
    const lobbyStore = useLobbyStore.getState()
    const { slots } = m
    lobbyStore.updateSlots(slots)
  }

  private onChatMessage = (m: TChatMsg) => {
    const chatStore = useChatStore.getState()
    const me = useAuthStore.getState().user
    const { message, sender } = m
    const isMe = sender?.id === me?.id
    const username = !sender ? "unknown: " : isMe ? "" : `${sender.username}: `
    const msg = `${username}${message}`
    chatStore.addMessage(msg)
  }

  private onConnectMessage = (msg: TConnMsg) => {
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

export interface ILobbyUser {
  id: string
  username: string
}

export enum EMsgType {
  Connection = "connection",
  Chat = "chat",
  Update = "update",
  MoveSlot = "move_slot_action",
}
export type TWsBaseMsg = { type: EMsgType}
export type TWsMessage = TConnMsg | TChatMsg | TUpdateMsg

export type TConnMsg = TWsBaseMsg & {
  isConnected: boolean
  user: ILobbyUser
}
const isConnectionMsg = (message: TWsMessage): message is TConnMsg => message.type === EMsgType.Connection

export type TChatMsg = TWsBaseMsg & {
  sender: ILobbyUser
  message: string
  isSystem: boolean
}
const isChatMsg = (message: TWsMessage): message is TChatMsg => message.type === EMsgType.Chat

export type TUpdateMsg = TWsBaseMsg & {
  slots: ILobbySlot[]
}
const isUpdateMsg = (message: TWsMessage): message is TUpdateMsg => message.type === EMsgType.Update
