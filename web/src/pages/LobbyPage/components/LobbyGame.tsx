import { ESuit } from "@/models/ICard"
import { EGameStage } from "@/models/IGame"
import { gameService } from "@/services/game.service"
import { Text } from "@/shared/ui-kit/Text"
import { useAuthStore } from "@/stores/auth.store"
import { useGameStore } from "@/stores/game.store"
import cn from "classnames"

export const LobbyGame = () => {
  const game = useGameStore((state) => state.game)
  const getPlayerName = useGameStore((state) => state.getPlayerName)

  return (
    <div className="flex flex-col flex-grow">
      <PlayersCards />
      <div className="flex flex-col w-full py-4">
        <GameStage />
        <div className="flex flex-col">
          {game?.round.trump && <Text>Козырь: {game.round.trump}</Text>}
          {game?.round.firstStepPlayerId && <Text>Ходит: {getPlayerName(game?.round.firstStepPlayerId)}</Text>}
          {game?.round.praiserId && <Text>Хвалит: {getPlayerName(game?.round.praiserId)}</Text>}
        </div>
      </div>
      <PraisingWindow />
    </div>
  )
}

const PraisingWindow = () => {
  const stage = useGameStore((state) => state.game?.stage)
  const praiserId = useGameStore((state) => state.game?.round.praiserId)
  const me = useAuthStore((state) => state.user)

  if (stage !== EGameStage.Praising) return null
  if (praiserId !== me?.id) return null

  const handlePraise = (suit: ESuit) => {
    gameService.praiseTrump(suit)
  }

  return (
    <div className="flex flex-col fixed z-10 top-[50%] left-[50%] -translate-x-[50%] -translate-y-[50%] rounded-xl bg-white text-black px-8 py-4">
      <Text type="subheader" className="text-center">Вы хвалите козырь</Text>
      <div className="flex flex-row flex-nowrap gap-4 justify-center py-4">
        <Text
          onClick={() => handlePraise(ESuit.Booby)}
          type="header" className={cn(
          "text-stopRed text-[42px] rounded-md border-2 border-s-slate-300 py-4 px-2",
          "cursor-pointer hover:scale-105"
        )}>
          {ESuit.Booby}
        </Text>
        <Text
          onClick={() => handlePraise(ESuit.Chervy)}
          type="header" className={cn(
          "text-stopRed text-[42px] rounded-md border-2 border-s-slate-300 py-4 px-2",
          "cursor-pointer hover:scale-105"
        )}>
          {ESuit.Chervy}
        </Text>
        <Text
          onClick={() => handlePraise(ESuit.Picky)}
          type="header" className={cn(
          "text-[42px] rounded-md border-2 border-s-slate-300 py-4 px-2",
          "cursor-pointer hover:scale-105"
        )}>
          {ESuit.Picky}
        </Text>
        <Text
          onClick={() => handlePraise(ESuit.Tref)}
          type="header" className={cn(
          "text-[42px] rounded-md border-2 border-s-slate-300 py-4 px-2",
          "cursor-pointer hover:scale-105"
        )}>
          {ESuit.Tref}
        </Text>
      </div>
    </div>
  )
}

const GameStage = () => {
  const stage = useGameStore((state) => state.game?.stage)

  const getStageText = () => {
    switch (stage) {
      case EGameStage.Preparing:
        return "Подготовка игры..."
      case EGameStage.Praising:
        return "Игрок хвалит козырь"
      case EGameStage.PlayerStep:
        return "Ход игрока"
      case EGameStage.DealerStep:
        return "Раздача карт"
      case EGameStage.End:
        return "Игра окончена"
      default:
        return ""
    }
  }

  if (stage === null || stage === undefined) return null
  return (
    <Text>Этап: {getStageText()}</Text>
  )
}

const PlayersCards = () => {
  const game = useGameStore((state) => state.game)
  const getPlayerName = useGameStore((state) => state.getPlayerName)

  return (
    <div className="flex flex-col gap-4">
      {game?.players.map((player) => (
        <div key={player.id} className="flex flex-col gap-1">
          {getPlayerName(player.id)}
          <div className="flex gap-2 p-1">
            {player.hand.map((card, i) => {
              if (!card || card.isHidden) return (
                <span
                  key={`hidden-${i}`}
                  className={cn(
                    "min-h-12 min-w-8 bg-hint",
                    "inline-flex justify-center text-sm px-1 py-3 rounded-md border-[1px] border-slate-900 cursor-default",
                  )}
                />
              )
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
  )
}