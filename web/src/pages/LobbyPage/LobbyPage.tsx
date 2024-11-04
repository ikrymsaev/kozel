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
import cn from "classnames"
import { ESuit } from "@/models/ICard"

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
    if (game) return <Game />

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

const Game = () => {
  const game = useGameStore((state) => state.game)

  console.log(game?.players)

  return (
    <div className="flex flex-col flex-grow">
      <div className="flex flex-col gap-4">
        {game?.players.map((player) => (
          <div key={player.id} className="flex flex-col gap-1">
            {player.name}
            <div className="flex gap-2 p-1">
              {player.hand.map((card) => {
                if (!card) return <div/>
                return (
                  <span
                    key={`${card.suit}${card.type}`}
                    className={cn(
                      (card.suit === ESuit.Booby || card.suit === ESuit.Chervy) ? "text-stopRed" : "text-black",
                      "min-h-12 min-w-8",
                      "inline-flex justify-center text-sm px-1 py-3 rounded-md border-[1px] border-slate-900 cursor-default hover:scale-105",
                      card.isHidden ? "bg-hint" : card.isTrump ?  "bg-yellow-200" : "bg-white",
                    )}
                  >
                    {card.type}{card.suit}
                  </span>
                )
              })}
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}