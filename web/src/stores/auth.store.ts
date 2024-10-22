import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { IPlayer } from "../models/IPlayer";

const getMockUser = (): IPlayer => {
  const mocked = localStorage.getItem('gk_user')
  if (mocked) return JSON.parse(mocked) as IPlayer

  const user = { id: 'i423ub6234iu6b', name: 'John Doe' }
  localStorage.setItem('gk_user', JSON.stringify(user))

  return { id: 'i423ub6234iu6b', name: 'John Doe' }
}

interface State {
  user: IPlayer | null
}
interface Actions {
  setUser: (user: IPlayer | null) => void
}
type TAuthStore = State & Actions

export const useAuthStore = create<TAuthStore>()(immer((set) => ({
  /**
   * To mock a user set it in local storage
   * If user is not set, it will be default
   */
  user: getMockUser(), //! REMOVE LATER
  setUser: (user: IPlayer | null) => set({ user }),
})))