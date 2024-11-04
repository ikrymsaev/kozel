import { ESuit, ICard } from "./ICard"

export interface IRound {
  firstStepPlayerId: string
  praiserId: string
  trump: ESuit
  bribes: [ICard[], ICard[]]
}