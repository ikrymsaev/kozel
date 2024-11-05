import { useGameStore } from "@/stores/game.store"
import { EWSAction } from "../api/ws/ws.actions"
import { EWSMessage, TCardActionMsg, TChangeTurnMsg, TNewTrumpMsg, TRoundResultMsg, TStageMsg, TStakeResultMsg, TUpdateGameStateMsg } from "../api/ws/ws.messages"
import { TWsService, wsService } from "../api/ws/ws.service"
import { ESuit, ICard } from "@/models/ICard"
import { toast } from "react-toastify"

class GameService {
  private readonly ws: TWsService
  constructor() {
    this.ws = wsService
    this.ws.listen(EWSMessage.UpdateGameState, this.onUpdateGameState)
    this.ws.listen(EWSMessage.Stage, this.onUpdateStage)
    this.ws.listen(EWSMessage.NewTrump, this.onNewTrump)
    this.ws.listen(EWSMessage.ChangeTurn, this.onChangeTurn)
    this.ws.listen(EWSMessage.CardAction, this.onCardAction)
    this.ws.listen(EWSMessage.StakeResult, this.onStakeResult)
    this.ws.listen(EWSMessage.RoundResult, this.onRoundResult)
  }

  public startGame = () => {
    this.ws.send({ type: EWSAction.StartGame })
  }

  public praiseTrump = (trump: ESuit) => {
    this.ws.send({ type: EWSAction.PraiseTrump, trump })
  }

  public pickCard = (card: ICard) => {
    this.ws.send({ type: EWSAction.MoveCard, cardId: card.id })
  }

  private onRoundResult = (msg: TRoundResultMsg) => {
    const gameStore = useGameStore.getState()
    toast(`Team ${msg.result.winnerTeam} win!!! Get +${msg.result.score} points!`)
    gameStore.setRoundResult(msg.result.winnerTeam, msg.result.score)
  }

  private onStakeResult = (msg: TStakeResultMsg) => {
    const gameStore = useGameStore.getState()
    toast(`${gameStore.getPlayerName(msg.result.winnerId)} get bribe with ${msg.result.bribeScore} points!`)
    gameStore.moveFromTableToBribes(msg.result.winnerId, msg.result.bribeScore)
  }

  private onCardAction = (msg: TCardActionMsg) => {
    const gameStore = useGameStore.getState()
    gameStore.replaceToTable(msg.card)
  }

  private onChangeTurn = (msg: TChangeTurnMsg) => {
    const gameStore = useGameStore.getState()
    gameStore.changeTurn(msg.turnPlayerId)
  }

  private onNewTrump = (msg: TNewTrumpMsg) => {
    const gameStore = useGameStore.getState()
    gameStore.setTrump(msg.trump)
    toast(`Козырь : ${msg.trump}`)
  }

  private onUpdateGameState = (msg: TUpdateGameStateMsg) => {
    const gameStore = useGameStore.getState()
    gameStore.updateGame(msg.game)
  }

  private onUpdateStage = (msg: TStageMsg) => {
    const gameStore = useGameStore.getState()
    gameStore.updateStage(msg.stage)
  }
}

export const gameService = new GameService()