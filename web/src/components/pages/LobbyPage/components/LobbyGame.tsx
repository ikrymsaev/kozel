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
          {game?.round.trump && <Text>Козырь: <Text type="header">{game.round.trump}</Text></Text>}
          {game?.round.firstStepPlayerId && <Text>Первый ход у {getPlayerName(game.round.firstStepPlayerId)}</Text>}
          {game?.round.turnPlayerId && <Text>Ходит: {getPlayerName(game.round.turnPlayerId)}</Text>}
          {game?.round.praiserId && <Text>Хвалит: {getPlayerName(game.round.praiserId)}</Text>}
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
      case EGameStage.GameOver:
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
  const table = useGameStore((state) => state.table)
  const getPlayerName = useGameStore((state) => state.getPlayerName)

  return (
    <div className="flex flex-col-reverse md:flex-row gap-4">
      <div className="flex flex-col gap-4 flex-1">
        {game?.players.map((player) => (
          <div key={player.id} className={cn(
            "flex flex-col gap-1 min-h-28 md:px-4 md:py-2 justify-center rounded-md text-black font-semibold",
            player.team === 1 && "bg-sky-400",
            player.team === 2 && "bg-red-400",
          )}>
            <div className="flex items-center gap-2">
              {player.id === game?.round.praiserId && <Text className="text-white font-bold">{">$<"}</Text>}
              {getPlayerName(player.id)}
              {game?.round.turnPlayerId === player.id && (
                <Text type="sm-1" className="text-white font-bold">{" <<< ход"}</Text>
              )}
            </div>
            <div className="flex gap-2 p-1">
              {player.hand.map((card) => {
                if (!card || card.isHidden) return (
                  <span
                    key={card.id}
                    className={cn(
                      "min-h-12 min-w-8 bg-gray-400",
                      "inline-flex justify-center text-sm px-1 py-3 rounded-md border-[1px] border-slate-900 cursor-default",
                    )}
                  />
                )
                return (
                  <span
                    key={card.id}
                    onClick={() => gameService.pickCard(card)}
                    className={cn(
                      (card.suit === ESuit.Booby || card.suit === ESuit.Chervy) ? "text-stopRed" : "text-black",
                      "min-h-12 min-w-8",
                      "inline-flex justify-center text-sm px-1 py-3 rounded-md border-[1px] border-slate-900 hover:scale-105",
                      card.isHidden ? "bg-gray-300 cursor-default" : card.isTrump ?  "bg-yellow-200 cursor-pointer" : "bg-white cursor-pointer",
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
      <div className="flex flex-col flex-1 border-2 border-s-slate-300 rounded-lg overflow-hidden">
        <Text className="text-center">Стол</Text>
        <div className="flex flex-col flex-1 items-center justify-center gap-2 p-2">
          <div className="flex flex-row items-start justify-center gap-2">
            {table.map((card) => {
              return (
                <span
                  key={card.id}
                  className={cn(
                    (card.suit === ESuit.Booby || card.suit === ESuit.Chervy) ? "text-stopRed" : "text-black",
                    "min-h-12 min-w-8",
                    "inline-flex justify-center text-sm px-1 py-3 rounded-md border-[1px] border-slate-900 cursor-default",
                    "bg-white",
                  )}
                >
                  {card.type}{card.suit}
                </span>
              )
            })}
          </div>
        </div>
        <Text className="text-center">Взятки</Text>
        <div className="flex flex-row flex-nowrap">
          <div className="flex flex-1 justify-center py-2 bg-sky-400">
            <Text>{game?.round.bribes[0]}</Text>
          </div>
          <div className="flex flex-1 justify-center py-2 bg-red-400">
            <Text>{game?.round.bribes[1]}</Text>
          </div>
        </div>
        <Text className="text-center mt-2">Общий счет</Text>
        <div className="flex flex-row flex-nowrap">
          <div className="flex flex-1 justify-center py-2 bg-sky-400">
            <Text>{game?.score[0]}</Text>
          </div>
          <div className="flex flex-1 justify-center py-2 bg-red-400">
            <Text>{game?.score[1]}</Text>
          </div>
        </div>
      </div>
    </div>
  )
}