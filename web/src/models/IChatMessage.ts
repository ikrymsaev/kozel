import { ILobbyMember } from "./IPlayer"

export interface IChatMessage {
  sender: ILobbyMember
  message: string
  isSystem: boolean
}