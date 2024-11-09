import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { toast } from "react-toastify";
import axios from "axios";
import { IAuth } from "@/models/IAuth";
import { getApiUrl } from "@/shared/utils/get-api-url";

interface State {
  user: IAuth | null
  token: string
  loading: boolean
}
interface Actions {
  signUp: (dto: TSignDto) => Promise<void>
  signIn: (dto: TSignDto) => Promise<void>
  signOut: () => void
}
type TAuthStore = State & Actions

type TSignDto = {
  username: string
  password: string
}

export const useAuthStore = create<TAuthStore>()(immer((set) => ({
  /**
   * *State
   */
  user: localStorage.getItem('gk_user') ? JSON.parse(localStorage.getItem('gk_user')!) : null,
  token: localStorage.getItem('gk_token') || '',
  loading: false,
  /**
   * *Actions
   */
  signUp: async (dto: TSignDto) => {
    set({ loading: true })
    try {
      const { data } = await axios
        .post<TSignDto, { data: { user: IAuth, token: string }}>(
          getApiUrl() + '/api/auth/signUp',
          dto
        )
      const { user, token } = data
      if (!user || !token) {
        toast('Failed to sign in', { type: 'error' })
        return
      }
      set({ user, token })
      localStorage.setItem('token', token)
      localStorage.setItem('user', JSON.stringify(user))
      toast('Successfully signed up', { type: 'success' })
    }
    catch (e) {
      console.error(e)
      toast('Failed to sign up', { type: 'error' })
    }
    finally {
      set({ loading: false })
    }
  },
  signIn: async (dto: TSignDto) => {
    set({ loading: false })
    try {
      const { data } = await axios
        .post<TSignDto, { data: { user: IAuth, token: string } }>(
          getApiUrl() + '/api/auth/signIn',
          dto
        )
      const { user, token } = data
      if (!user || !token) {
        toast('Failed to sign in', { type: 'error' })
        return
      }
      set({ user, token })
      localStorage.setItem('gk_token', token)
      localStorage.setItem('gk_user', JSON.stringify(user))
      toast('Successfully signed ', { type: 'success' })
    }
    catch (e) {
      console.error(e)
      toast('Failed to sign in', { type: 'error' })
    }
    finally {
      set({ loading: false })
    }
  },
  signOut: () => {
    set({ user: null, token: '' });
    localStorage.removeItem('gk_token')
    localStorage.removeItem('gk_user')
    toast('Signed out', { type: 'success' })
  }
})))