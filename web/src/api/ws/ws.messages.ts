/**
 * @Messages received from the server
 */

import { ESuit, ICard } from "@/models/ICard"
import { EGameStage, IGameState } from "@/models/IGame"
import { ILobbySlot } from "@/models/ILobby"
import { ILobbyMember } from "@/models/IPlayer"

export enum EWSMessage {
  Error,
  Connection,
  NewMessage,
  UpdateSlots,
  UpdateGameState,
  Stage,
  NewTrump,
  ChangeTurn,
  CardAction
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
/** Стадия игры */
export type TStageMsg = TWsBaseMsg & {
  stage: EGameStage
}
/** Новый козырь */
export type TNewTrumpMsg = TWsBaseMsg & {
  trump: ESuit
}

export type TChangeTurnMsg = TWsBaseMsg & {
  turnPlayerId: string
}

export type TCardActionMsg = TWsBaseMsg & {
  playerId: string
  card: ICard
}

export type TWsMessage =
  | TErrorMsg
  | TConnectionMsg
  | TNewMessageMsg
  | TUpdateSlotsMsg
  | TUpdateGameStateMsg
  | TStageMsg
  | TNewTrumpMsg
  | TChangeTurnMsg
  | TCardActionMsg

export type TMsgMap = {
  [EWSMessage.Error]: TErrorMsg
  [EWSMessage.Connection]: TConnectionMsg
  [EWSMessage.NewMessage]: TNewMessageMsg
  [EWSMessage.UpdateSlots]: TUpdateSlotsMsg
  [EWSMessage.UpdateGameState]: TUpdateGameStateMsg
  [EWSMessage.Stage]: TStageMsg
  [EWSMessage.NewTrump]: TNewTrumpMsg
  [EWSMessage.ChangeTurn]: TChangeTurnMsg
  [EWSMessage.CardAction]: TCardActionMsg
}

export const isWsMsg = <T extends TWsMessage>(msg: TWsMessage, type: EWSMessage): msg is T => msg.type === type