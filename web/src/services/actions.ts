/**
 * @Actions to be sent to the server
 */

export enum EWSAction {
  SendMessage,
  MoveSlot,
  StartGame
}

export type TWsBaseAction = { type: EWSAction }

/** Отправить сообщение в чат */
export type TSendMessage = TWsBaseAction & {
  message: string
}
/** Переместить слот */
export type TMoveSlot = TWsBaseAction & {
  from: number,
  to: number
}
/** Начать игру */
export type TStartGame = TWsBaseAction

export type TWSAction =
  | TSendMessage
  | TMoveSlot
  | TStartGame