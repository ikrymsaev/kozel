import { ICard } from "./ICard"

export interface IUser {
  id: string
  username: string
}

export interface IPlayer {
	id: string
	name: string
  hand: ICard[]
  position: number
  user: IUser
  team: number
}

/** Информация о члене лобби */
export interface ILobbyMember {
  id: string
  username: string
}