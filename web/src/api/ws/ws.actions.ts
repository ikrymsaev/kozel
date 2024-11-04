/**
 * @Actions to be sent to the server
 */

import { ESuit } from "@/models/ICard"

export enum EWSAction {
  SendMessage,
  MoveSlot,
  StartGame,
  PraiseTrump
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
/** Захвалить козырь. */
export type TPraiseTrump = TWsBaseAction & {
  trump: ESuit
}


export type TWSAction =
  | TSendMessage
  | TMoveSlot
  | TStartGame
  | TPraiseTrump