import { IGameState } from "@/models/IGame";
import { create } from "zustand";
import { immer } from "zustand/middleware/immer";


interface State {
  game: IGameState | null
}
interface Actions {
  updateGame: (game: IGameState) => void
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