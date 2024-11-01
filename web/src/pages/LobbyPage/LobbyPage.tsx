import { MainLayout } from "../layouts/MainLayout/MainLayout"
import { useEffect } from "react"
import { useSettingsStore } from "../../stores/settings.store"
import { useSearchParams } from "react-router-dom"
import { lobbyService } from "../../services/lobby.service"
import { LobbySlots } from "./components/LobbySlots"
import { LobbyChat } from "./components/LobbyChat"
import { Button } from "@/shared/ui-kit/Button"

export const LobbyPage = () => {
  const [searchParams] = useSearchParams();
  const lobbyId = searchParams.get("id");

  const fetchDeck = useSettingsStore((state) => state.fetchDeck)

  useEffect(() => {
    fetchDeck()
  }, [fetchDeck])

  /** Connect to websocket chat */
  useEffect(() => {
    if (!lobbyId) return
    lobbyService.joinLobby(lobbyId)
    return lobbyService.leaveLobby
  }, [lobbyId])

  return (
    <MainLayout>
      <h1>Lobby Page</h1>
      <LobbyChat />
      <LobbySlots />
      <Button size="m" color="gold" onClick={lobbyService.startGame} >Start Game</Button>
    </MainLayout>
  )
}