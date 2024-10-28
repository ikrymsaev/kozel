import { create } from "zustand";
import { immer } from "zustand/middleware/immer";


interface State {
  messages: string[]
}
interface Actions {
  addMessage: (msg: string) => void
}
type TChatStore = State & Actions;

export const useChatStore = create<TChatStore>()(immer((set) => ({
  /** State */
  messages: [],
  /** Actions */
  addMessage: (msg: string) => {
    set((state) => {
      state.messages.push(msg)
    })
  },
})))