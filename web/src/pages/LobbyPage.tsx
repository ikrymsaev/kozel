import { useSearchParams } from "react-router-dom"
import { MainLayout } from "./layouts/MainLayout/MainLayout"
import { useEffect } from "react"
import { wsService } from "../services/ws.service"

export const LobbyPage = () => {
  const [params] = useSearchParams()
  const id = params.get("id")

  /**
   * Test connect to websocket chat
   * http://localhost:5173/game?userId=2&username=zach from one browser
   * http://localhost:5173/game?userId=1&username=hellfever from another
   */
  useEffect(() => {
    const userId = id;
    const username = "hellfever" 
    if (!id || !userId || !username) return
    wsService.connect({ roomId: id, userId, username })
  }, [id])

  return (
    <MainLayout>
      <h1>Lobby Page</h1>
    </MainLayout>
  )
}