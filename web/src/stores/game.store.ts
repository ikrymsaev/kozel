import { ESuit } from "@/models/ICard";
import { EGameStage, IGameState } from "@/models/IGame";
import { create } from "zustand";
import { immer } from "zustand/middleware/immer";


interface State {
  game: IGameState | null
}
interface Actions {
  updateGame: (game: IGameState) => void
  updateStage: (stage: EGameStage) => void
  setTrump: (trump: ESuit) => void
  reset: () => void
  getPlayerName: (playerId: string) => string
}
type TGameStore = State & Actions;

export const useGameStore = create<TGameStore>()(immer((set, get) => ({
  /** State */
  game: null,
  /** Actions */
  updateGame: (game: IGameState) => set((state) => {
    state.game = game
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
    game: null
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