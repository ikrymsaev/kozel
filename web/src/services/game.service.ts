import { useGameStore } from "@/stores/game.store"
import { EWSAction } from "../api/ws/ws.actions"
import { EWSMessage, TNewTrumpMsg, TStageMsg, TUpdateGameStateMsg } from "../api/ws/ws.messages"
import { TWsService, wsService } from "../api/ws/ws.service"
import { ESuit } from "@/models/ICard"

class GameService {
  private readonly ws: TWsService
  constructor() {
    this.ws = wsService
    this.ws.listen(EWSMessage.UpdateGameState, this.onUpdateGameState)
    this.ws.listen(EWSMessage.Stage, this.onUpdateStage)
    this.ws.listen(EWSMessage.NewTrump, this.onNewTrump)
  }

  public startGame = () => {
    this.ws.send({ type: EWSAction.StartGame })
  }

  public praiseTrump = (trump: ESuit) => {
    this.ws.send({ type: EWSAction.PraiseTrump, trump })
  }

  private onNewTrump = (msg: TNewTrumpMsg) => {
    const gameStore = useGameStore.getState()
    gameStore.setTrump(msg.trump)
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