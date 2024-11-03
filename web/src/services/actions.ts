/**
 * @Actions to be sent to the server
 */

export enum EWSAction {
  SendMessage,
  MoveSlot,
}

export type TWsBaseAction = { type: EWSAction }

/** Отправить сообщение в чат */
export type TSendMessage = {
  message: string
}
/** Переместить слот */
export type TMoveSlot = {
  from: number,
  to: number
}

export type TWSAction =
  | TSendMessage
  | TMoveSlot

export type TActionsMap = {
  [EWSAction.SendMessage]: TSendMessage
  [EWSAction.MoveSlot]: TMoveSlot
}