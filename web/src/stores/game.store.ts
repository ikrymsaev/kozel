import { IGameState } from "@/models/IGame";
import { create } from "zustand";
import { immer } from "zustand/middleware/immer";


interface State {
  game: IGameState | null
}
interface Actions {
  updateGame: (game: IGameState) => void
  reset: () => void
}
type TGameStore = State & Actions;

export const useGameStore = create<TGameStore>()(immer((set) => ({
  /** State */
  game: null,
  /** Actions */
  updateGame: (game: IGameState) => set((state) => {
    state.game = game
  }),
  reset: () => set(() => ({
    game: null
  })),
})))