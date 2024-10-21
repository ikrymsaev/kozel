import { IPlayer } from "./IPlayer"

export interface IGame {
  id: string
  name: string
  players: IPlayer[]
}