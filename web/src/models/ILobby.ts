import { IPlayer } from "./IPlayer"

export interface ILobby {
  id: string
  name: string
  players: IPlayer[]
}

export interface ILobbySlot {
  order: number
  player?: IPlayer
}