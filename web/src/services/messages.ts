/**
 * @Messages received from the server
 */

import { IGameState } from "@/models/IGame"
import { ILobbySlot } from "@/models/ILobby"
import { ILobbyMember } from "@/models/IPlayer"

export enum EWSMessage {
  Error,
  Connection,
  NewMessage,
  UpdateSlots,
  UpdateGameState
}

export type TWsBaseMsg = { type: EWSMessage}

/** Ошибка */
export type TErrorMsg = TWsBaseMsg & {
  error: string
}
/** Кто-то подключился\отключился */
export type TConnectionMsg = TWsBaseMsg & {
  isConnected: boolean,
  user: ILobbyMember
}
/** Новое сообщение в чате */
export type TNewMessageMsg = TWsBaseMsg & {
  sender: ILobbyMember,
  message: string,
  isSystem: boolean
}
/** Слоты в лобби обновились */
export type TUpdateSlotsMsg = TWsBaseMsg & {
  slots: ILobbySlot[]
}
/** Обновилось состояние игры */
export type TUpdateGameStateMsg = TWsBaseMsg & {
  game: IGameState
}

export type TWsMessage =
  | TErrorMsg
  | TConnectionMsg
  | TNewMessageMsg
  | TUpdateSlotsMsg
  | TUpdateGameStateMsg

export type TMsgMap = {
  [EWSMessage.Error]: TErrorMsg
  [EWSMessage.Connection]: TConnectionMsg
  [EWSMessage.NewMessage]: TNewMessageMsg
  [EWSMessage.UpdateSlots]: TUpdateSlotsMsg
  [EWSMessage.UpdateGameState]: TUpdateGameStateMsg
}

export const isWsMsg = <T extends TWsMessage>(msg: TWsMessage, type: EWSMessage): msg is T => msg.type === type