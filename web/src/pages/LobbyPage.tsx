import { MainLayout } from "./layouts/MainLayout/MainLayout"
import { useEffect } from "react"
import { wsService } from "../services/ws.service"
import { useAuthStore } from "../stores/auth.store"

export const LobbyPage = () => {
  const user = useAuthStore((state) => state.user)

  /**
   * Test connect to websocket chat
   */
  useEffect(() => {
    if (!user) return
    const { id: userId, name: username } = user;
    wsService.connect({ roomId: userId, userId, username })
  }, [user])

  return (
    <MainLayout>
      <h1>Lobby Page</h1>
    </MainLayout>
  )
}