import { TWSAction, TWsBaseAction } from "./actions"
import { EWSMessage, isWsMsg, TConnectionMsg, TMsgMap, TNewMessageMsg, TUpdateSlotsMsg, TWsMessage } from "./messages"
import { WS } from "./ws"

type Listeners = {
  [key in EWSMessage]: Set<(data: TMsgMap[key]) => void>
}

class WSService extends WS {
  constructor() {
    super()
  }

  private readonly listeners: Listeners = {
    [EWSMessage.Connection]: new Set(),
    [EWSMessage.NewMessage]: new Set(),
    [EWSMessage.UpdateSlots]: new Set(),
    [EWSMessage.Error]: new Set(),
  }

  public connect(params: string) {
    super.conn(params)
    this.withConn((conn) => (conn.onmessage = this.onMessage))
  }

  public disconnect() {
    super.disconn()
  }

  public send = <T extends TWSAction>(data: T & TWsBaseAction) => {
    this.withConn((conn) => conn.send(JSON.stringify(data)))
  }

  public listen<T extends EWSMessage>(messageType: T, callback: (data: TMsgMap[T]) => void) {
    this.listeners[messageType].add(callback)
    return () => this.listeners[messageType].delete(callback)
  }

  private onMessage = (e: { data: string }) => {
    try {
      const message = JSON.parse(e.data) as TWsMessage
      console.log('WS: ', message)
      if (isWsMsg<TConnectionMsg>(message, EWSMessage.Connection))
        return this.listeners[EWSMessage.Connection].forEach((cb) => cb(message))
      if (isWsMsg<TNewMessageMsg>(message, EWSMessage.NewMessage))
        return this.listeners[EWSMessage.NewMessage].forEach((cb) => cb(message))
      if (isWsMsg<TUpdateSlotsMsg>(message, EWSMessage.UpdateSlots))
        return this.listeners[EWSMessage.UpdateSlots].forEach((cb) => cb(message))
      console.error('Unknown message: ', message)
    } catch (e) {
      console.error('Failed to parse message: ', e)
    }
  }
}

export type TWsService = WSService
export const wsService = new WSService()