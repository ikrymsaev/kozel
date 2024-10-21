import { TConnectionParams, WS } from "./ws"

type DataTypes = {
  [EMessageType.System]: { content: string }
  [EMessageType.GameEvent]: TGameEvent
}

type Listeners = {
  [key in EMessageType]: Set<(data: DataTypes[key]) => void>
}

class WSService extends WS {
  constructor() {
    super()
  }

  private readonly listeners: Listeners = {
    [EMessageType.System]: new Set(),
    [EMessageType.GameEvent]: new Set()
  }

  public connect(params: TConnectionParams) {
    super.conn(params)
    this.withConn((conn) => (conn.onmessage = this.onMessage))
  }

  private onMessage = (e: { data: string }) => {
    try {
      const message = JSON.parse(e.data) as TWsMessage
      console.log(message)
      if (isSystemMessage(message))
        return this.listeners[EMessageType.System].forEach((cb) => cb(message.body))
      if (isGameEvent(message))
        return this.listeners[EMessageType.GameEvent].forEach((cb) => cb(message.body))
      console.error('Unknown message: ', message)
    } catch (e) {
      console.error('Failed to parse message: ', e)
    }
  }

  public send = (data: TWsMessage) => {
    this.withConn((conn) => conn.send(JSON.stringify(data)))
  }

  public listen<T extends EMessageType>(messageType: T, callback: (data: DataTypes[T]) => void) {
    this.listeners[messageType].add(callback)
    return () => this.listeners[messageType].delete(callback)
  }
}

export const wsService = new WSService()

export enum EMessageType {
  System,
  GameEvent
}

export type TWsMessage = TSystemMsg | TGameMsg
export type TWsBaseMessage = { info: { messageType: EMessageType; roomId: string } }
export type TSystemMsg = TWsBaseMessage & { body: { content: string } }
export type TGameMsg = TWsBaseMessage & { body: TGameEvent }

/** Game event types */
enum EGameEvent {
  UserAction,
  Alert
}
type TUserAction = { userId: string; action: string }
type TAlert = { content: string }
type TGameEventData<T extends EGameEvent> = T extends EGameEvent.UserAction
  ? TUserAction
  : T extends EGameEvent.Alert
    ? TAlert
    : never
type TGameEvent = { gameId: string; event: EGameEvent; data: TGameEventData<EGameEvent> }

const isSystemMessage = (message: TWsMessage): message is TSystemMsg =>
  message.info.messageType === EMessageType.System
const isGameEvent = (message: TWsMessage): message is TSystemMsg =>
  message.info.messageType === EMessageType.GameEvent
