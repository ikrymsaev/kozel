import { ESuit } from "./ICard"

export interface IRound {
  firstStepPlayerId: string
  turnPlayerId: string
  praiserId: string
  trump: ESuit
  bribes: [number, number]
}