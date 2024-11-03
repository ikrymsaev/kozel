import { MainLayout } from "../layouts/MainLayout/MainLayout"
import { useEffect } from "react"
import { useSettingsStore } from "../../stores/settings.store"
import { useSearchParams } from "react-router-dom"
import { lobbyService } from "../../services/lobby.service"
import { LobbySlots } from "./components/LobbySlots"
import { LobbyChat } from "./components/LobbyChat"
import { Button } from "@/shared/ui-kit/Button"
import { gameService } from "@/services/game.service"

export const LobbyPage = () => {
  const [searchParams] = useSearchParams();
  const lobbyId = searchParams.get("id");

  const fetchDeck = useSettingsStore((state) => state.fetchDeck)

  useEffect(() => {
    fetchDeck()
  }, [fetchDeck])

  /** Connect to websocket */
  useEffect(() => {
    if (!lobbyId) return
    lobbyService.joinLobby(lobbyId)
    return lobbyService.leaveLobby
  }, [lobbyId])

  return (
    <MainLayout>
      <h1>Lobby Page</h1>
      <LobbySlots />
      <div className="flex w-full justify-center py-4">
        <Button size="m" color="gold"
          onClick={gameService.startGame}
        >
          Start Game
        </Button>
      </div>
      <LobbyChat />
    </MainLayout>
  )
}