import { useAuthStore } from "@/stores/auth.store"
import { PropsWithChildren, useEffect } from "react"
import { useNavigate } from "react-router-dom"

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const auth = useAuthStore((state) => state.user)
  const token = useAuthStore((state) => state.token)

  const navigate = useNavigate()

  useEffect(() => {
    if (!auth || !token) {
      navigate('/auth')
    }
  }, [auth, token, navigate])

  if (!auth || !token) {
    return null
  }
  return children
}