import { IPlayer } from "./IPlayer"

export interface ILobby {
  id: string
  name: string
  players: IPlayer[]
}