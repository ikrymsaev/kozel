import { MainLayout } from "../layouts/MainLayout/MainLayout"
import { useEffect } from "react"
import { useSettingsStore } from "../../stores/settings.store"
import { useSearchParams } from "react-router-dom"
import { lobbyService } from "../../services/lobby.service"
import { LobbySlots } from "./components/LobbySlots"
import { LobbyChat } from "./components/LobbyChat"
import { Button } from "@/shared/ui-kit/Button"
import { gameService } from "@/services/game.service"
import { useGameStore } from "@/stores/game.store"
import { LobbyGame } from "./components/LobbyGame"

export const LobbyPage = () => {
  const [searchParams] = useSearchParams();
  const lobbyId = searchParams.get("id");

  const fetchDeck = useSettingsStore((state) => state.fetchDeck)
  const game = useGameStore((state) => state.game)
  const resetGame = useGameStore((state) => state.reset)

  useEffect(() => {
    fetchDeck()
  }, [fetchDeck])

  /** Connect to websocket */
  useEffect(() => {
    if (!lobbyId) return
    lobbyService.joinLobby(lobbyId)
    return () => {
      lobbyService.leaveLobby()
      resetGame()
    }
  }, [lobbyId, resetGame])
  

  const PageContent = () => {
    if (game) return <LobbyGame />

    return (
      <div className="flex flex-col flex-grow">
        <LobbySlots />
        <div className="flex w-full justify-center py-4">
          <Button size="m" color="gold"
            onClick={gameService.startGame}
          >
            Start Game
          </Button>
        </div>
      </div>
    )
  }

  return (
    <MainLayout>
      <div className="flex flex-col flex-grow overflow-y-auto">
        <PageContent />
      </div>
      <LobbyChat />
    </MainLayout>
  )
}

