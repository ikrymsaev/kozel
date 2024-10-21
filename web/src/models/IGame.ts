import { IPlayer } from "./IPlayer"

export interface IGame {
  id: number
  name: string
  players: IPlayer[]
}