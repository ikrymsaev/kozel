import { useEffect } from "react"
import { wsService } from "../services/ws.service"
import { useSearchParams } from "react-router-dom"
import { MainLayout } from "./layouts/MainLayout/MainLayout"

export const GamePage = () => {
  const [params] = useSearchParams()
  const userId = params.get("userId")
  const username = params.get("username")

  /**
   * Test connect to websocket chat
   * http://localhost:5173/game?userId=2&username=zach from one browser
   * http://localhost:5173/game?userId=1&username=hellfever from another
   */
  useEffect(() => {
    console.log({userId, username})
    const roomId = "132"
    if (!roomId || !userId || !username) return
    wsService.connect({ roomId, userId, username })
  }, [userId, username])

  return (
    <MainLayout>
      <h1>Game Page</h1>
    </MainLayout>
  )
}