/**
 * @Messages received from the server
 */

import { ILobbySlot } from "@/models/ILobby"
import { ILobbyMember } from "@/models/IPlayer"

export enum EWSMessage {
  Error,
  Connection,
  NewMessage,
  UpdateSlots,
}

export type TWsBaseMsg = { type: EWSMessage}

export type TErrorMsg = TWsBaseMsg & {
  error: string
}
export type TConnectionMsg = TWsBaseMsg & {
  isConnected: boolean,
  user: ILobbyMember
}
export type TNewMessageMsg = TWsBaseMsg & {
  sender: ILobbyMember,
  message: string,
  isSystem: boolean
}
export type TUpdateSlotsMsg = TWsBaseMsg & {
  slots: ILobbySlot[]
}

export type TWsMessage =
  | TErrorMsg
  | TConnectionMsg
  | TNewMessageMsg
  | TUpdateSlotsMsg

export type TMsgMap = {
  [EWSMessage.Error]: TErrorMsg
  [EWSMessage.Connection]: TConnectionMsg
  [EWSMessage.NewMessage]: TNewMessageMsg
  [EWSMessage.UpdateSlots]: TUpdateSlotsMsg
}

export const isWsMsg = <T extends TWsMessage>(msg: TWsMessage, type: EWSMessage): msg is T => msg.type === type