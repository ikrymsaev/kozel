import { MainLayout } from "./layouts/MainLayout/MainLayout"
import { useEffect } from "react"
import { useSettingsStore } from "../stores/settings.store"

export const LobbyPage = () => {
  const fetchDeck = useSettingsStore((state) => state.fetchDeck)

  useEffect(() => {
    fetchDeck()
  }, [fetchDeck])

  /**
   * Test connect to websocket chat
   */
  // useEffect(() => {
  //   if (!user) return
  //   const { id: userId, name: username } = user;
  //   wsService.connect({ roomId: userId, userId, username })
  // }, [user])

  return (
    <MainLayout>
      <h1>Lobby Page</h1>
    </MainLayout>
  )
}