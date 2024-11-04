import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { IUser } from "../models/IPlayer";

const getMockUser = (): IUser => {
  const mocked = localStorage.getItem('gk_user')
  if (mocked) return JSON.parse(mocked) as IUser

  const user = { id: 'i423ub6234iu6b', username: 'John Doe' }
  localStorage.setItem('gk_user', JSON.stringify(user))

  return { id: 'i423ub6234iu6b', username: 'John Doe' }
}

interface State {
  user: IUser | null
}
interface Actions {
  setUser: (user: IUser | null) => void
}
type TAuthStore = State & Actions

export const useAuthStore = create<TAuthStore>()(immer((set) => ({
  /**
   * To mock a user set it in local storage
   * If user is not set, it will be default
   */
  user: getMockUser(), //! REMOVE LATER
  setUser: (user: IUser | null) => set({ user }),
})))