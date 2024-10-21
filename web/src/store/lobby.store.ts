import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { IGame } from "../models/IGame";

interface State {
  activeGames: IGame[]
}
interface Actions {
  setActiveGames: (activeGames: IGame[]) => void
  addGame: (game: IGame) => void
  removeGame: (id: IGame['id']) => void
}
type TLobbyStore = State & Actions;

export const useLobbyStore = create<TLobbyStore>()(immer((set) => ({
  /** State */
  activeGames: [],
  /** Actions */
  setActiveGames: (activeGames: IGame[]) => {set({ activeGames })},
  addGame: (game: IGame) => set((state) => {
    state.activeGames.push(game)
  }),
  removeGame: (id: IGame['id']) => set((state) => {
    state.activeGames = state.activeGames.filter((game) => game.id !== id)
  }),
})))