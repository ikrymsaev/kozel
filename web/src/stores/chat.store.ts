import { IChatMessage } from "@/models/IChatMessage";
import { create } from "zustand";
import { immer } from "zustand/middleware/immer";


interface State {
  messages: IChatMessage[]
}
interface Actions {
  addMessage: (msg: IChatMessage) => void
  reset: () => void
}
type TChatStore = State & Actions;

export const useChatStore = create<TChatStore>()(immer((set) => ({
  /** State */
  messages: [],
  /** Actions */
  addMessage: (msg: IChatMessage) => set((state) => {
    state.messages.push(msg)
  }),
  reset: () => set((state) => {
    state.messages = [];
  })
})))