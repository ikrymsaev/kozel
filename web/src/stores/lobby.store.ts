import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { ILobby, ILobbySlot } from "../models/ILobby";

interface State {
  activeGames: ILobby[]
  slots: ILobbySlot[]
}
interface Actions {
  setActiveGames: (activeGames: ILobby[]) => void
  addLobby: (game: ILobby) => void
  removeLobby: (id: ILobby['id']) => void
  updateSlots: (slots: ILobbySlot[]) => void
}
type TLobbyStore = State & Actions;

export const useLobbyStore = create<TLobbyStore>()(immer((set) => ({
  /**
   * *State
   */
  activeGames: [],
  slots: [],
  /**
   * *Actions
   */
  setActiveGames: (activeGames: ILobby[]) => {
    set({ activeGames })
  },
  addLobby: (game: ILobby) => set((state) => {
    state.activeGames.unshift(game)
  }),
  removeLobby: (id: ILobby['id']) => {
    set((state) => {
      state.activeGames = state.activeGames.filter((game) => game.id !== id)
    })
  },
  updateSlots: (slots: ILobbySlot[]) => set({ slots }),
})))