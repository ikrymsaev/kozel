
export interface ICard {
  suit: ESuit
  type: ECardType
  isTrump: boolean
}

/** Типы карт */
export enum ECardType {
  Seven = "7",
  Eight = "8",
  Nine = "9",
  Ten = "10",
  Jack = "J",
  Queen = "Q",
  King = "K",
  Ace = "A"
}
/** Масти карт */
export enum ESuit {
  Booby = "♦",
  Chervy = "♥",
  Picky = "♠",
  Tref = "♣"
}

export interface ICardType {
  type: ECardType
  name: string
  order: number
  score: number
}
export interface ISuit {
  suit: ESuit
  name: string
  order: number
}