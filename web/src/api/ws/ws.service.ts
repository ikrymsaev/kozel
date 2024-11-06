import { TWSAction } from "./ws.actions"
import { EWSMessage, isWsMsg, TCardActionMsg, TChangeTurnMsg, TConnectionMsg, TErrorMsg, TGameOverMsg, TMsgMap, TNewMessageMsg, TNewTrumpMsg, TRoundResultMsg, TStageMsg, TStakeResultMsg, TUpdateGameStateMsg, TUpdateSlotsMsg, TWsMessage } from "./ws.messages"
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
    [EWSMessage.UpdateGameState]: new Set(),
    [EWSMessage.Stage]: new Set(),
    [EWSMessage.NewTrump]: new Set(),
    [EWSMessage.ChangeTurn]: new Set(),
    [EWSMessage.CardAction]: new Set(),
    [EWSMessage.StakeResult]: new Set(),
    [EWSMessage.RoundResult]: new Set(),
    [EWSMessage.GameOver]: new Set(),
  }

  public connect(params: string) {
    super.conn(params)
    this.withConn((conn) => (conn.onmessage = this.onMessage))
  }

  public disconnect() {
    super.disconn()
  }

  public send = <T extends TWSAction>(data: T) => {
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
      /** Кто-то подключился\отключился */
      if (isWsMsg<TConnectionMsg>(message, EWSMessage.Connection)) {
        return this.listeners[EWSMessage.Connection].forEach((cb) => cb(message))
      }
      /** Новое сообщение в чате */
      if (isWsMsg<TNewMessageMsg>(message, EWSMessage.NewMessage)){
        return this.listeners[EWSMessage.NewMessage].forEach((cb) => cb(message))
      }
      /** Слоты в лобби обновились */
      if (isWsMsg<TUpdateSlotsMsg>(message, EWSMessage.UpdateSlots)) {
        return this.listeners[EWSMessage.UpdateSlots].forEach((cb) => cb(message))
      }
      /** Ошибка */
      if (isWsMsg<TErrorMsg>(message, EWSMessage.Error)) {
        return this.listeners[EWSMessage.Error].forEach((cb) => cb(message))
      }
      /** Обновилось состояние игры */
      if (isWsMsg<TUpdateGameStateMsg>(message, EWSMessage.UpdateGameState)) {
        return this.listeners[EWSMessage.UpdateGameState].forEach((cb) => cb(message))
      }
      /** Стадия игры изменилась */
      if (isWsMsg<TStageMsg>(message, EWSMessage.Stage)) {
        return this.listeners[EWSMessage.Stage].forEach((cb) => cb(message))
      }
      /** Новый козырь */
      if (isWsMsg<TNewTrumpMsg>(message, EWSMessage.NewTrump)) {
        return this.listeners[EWSMessage.NewTrump].forEach((cb) => cb(message))
      }
      /** Смена хода */
      if (isWsMsg<TChangeTurnMsg>(message, EWSMessage.ChangeTurn)) {
        return this.listeners[EWSMessage.ChangeTurn].forEach((cb) => cb(message))
      }
      /** Ход игрока */
      if (isWsMsg<TCardActionMsg>(message, EWSMessage.CardAction)) {
        return this.listeners[EWSMessage.CardAction].forEach((cb) => cb(message))
      }
      /** Результат кона */
      if (isWsMsg<TStakeResultMsg>(message, EWSMessage.StakeResult)) {
        return this.listeners[EWSMessage.StakeResult].forEach((cb) => cb(message))
      }
      /** Результат раунда */
      if (isWsMsg<TRoundResultMsg>(message, EWSMessage.RoundResult)) {
        return this.listeners[EWSMessage.RoundResult].forEach((cb) => cb(message))
      }
      /** Игра закончилась */
      if (isWsMsg<TGameOverMsg>(message, EWSMessage.GameOver)) {
        return this.listeners[EWSMessage.GameOver].forEach((cb) => cb(message))
      }
      console.error('Unknown message: ', message)
    } catch (e) {
      console.error('Failed to parse message: ', e)
    }
  }
}

export type TWsService = WSService
export const wsService = new WSService()