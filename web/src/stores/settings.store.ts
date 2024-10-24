import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { ICard } from "../models/ICard";

interface State {
  deck: ICard[] | null
}
interface Actions {
  fetchDeck: () => Promise<void>
}
type TSettingsStore = State & Actions;

export const useSettingsStore = create<TSettingsStore>()(immer((set) => ({
  /** State */
  deck: null,
  /** Actions */
  fetchDeck: async () => {
    try {
      const res = await fetch('http://localhost:8080/settings/deck')
      const deck = await res.json() || null
      set({ deck })
    } catch (e) {
      console.error(e)
    }
  },
})))