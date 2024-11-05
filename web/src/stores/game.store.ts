import { ESuit, ICard } from "@/models/ICard";
import { EGameStage, IGameState } from "@/models/IGame";
import { create } from "zustand";
import { immer } from "zustand/middleware/immer";


interface State {
  game: IGameState | null
  table: ICard[]
}
interface Actions {
  updateGame: (game: IGameState) => void
  changeTurn: (playerId: string) => void
  updateStage: (stage: EGameStage) => void
  setTrump: (trump: ESuit) => void
  reset: () => void
  getPlayerName: (playerId: string) => string
  replaceToTable: (card: ICard) => void
}
type TGameStore = State & Actions;

export const useGameStore = create<TGameStore>()(immer((set, get) => ({
  /** State */
  game: null,
  table: [],
  /** Actions */
  replaceToTable: (card: ICard) => set((state) => {
    if (!state.game) return
    for (const player of state.game.players) {
      player.hand = player.hand.filter((c) => c.id !== card.id)
    }
    state.table.push(card)
  }),
  updateGame: (game: IGameState) => set((state) => {
    state.game = game
  }),
  changeTurn: (playerId: string) => set((state) => {
    if (!state.game?.round) return
    state.game.round.turnPlayerId = playerId
  }),
  setTrump: (trump: ESuit) => set((state) => {
    if (!state.game?.round) return
    state.game.round.trump = trump
    for (const player of state.game.players) {
      for (const card of player.hand) {
        if (!card || card.isHidden) continue
        if (card.suit === trump) {
          card.isTrump = true
        }
      }
    }
  }),
  updateStage: (stage: EGameStage) => set((state) => {
    console.log("updateStage", stage)
    if (!state.game) return
    state.game.stage = stage
  }),
  reset: () => set(() => ({
    game: null,
    table: []
  })),
  getPlayerName: (playerId: string): string => {
    const game = get().game
    if (!game) return ''

    const player = game.players.find((player) => player.id === playerId)
    if (!player) return ''

    const botPostfix = !player.user ? ' (BOT)' : ''

    return player.name + botPostfix
  }
})))