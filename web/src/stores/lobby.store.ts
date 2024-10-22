import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { ILobby } from "../models/ILobby";

interface State {
  activeGames: ILobby[]
}
interface Actions {
  setActiveGames: (activeGames: ILobby[]) => void
  addLobby: (game: ILobby) => void
  removeGame: (id: ILobby['id']) => void
}
type TLobbyStore = State & Actions;

export const useLobbyStore = create<TLobbyStore>()(immer((set) => ({
  /** State */
  activeGames: [],
  /** Actions */
  setActiveGames: (activeGames: ILobby[]) => {set({ activeGames })},
  addLobby: (game: ILobby) => set((state) => {
    state.activeGames.push(game)
  }),
  removeGame: (id: ILobby['id']) => set((state) => {
    state.activeGames = state.activeGames.filter((game) => game.id !== id)
  }),
})))