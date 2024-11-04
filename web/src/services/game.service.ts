import { useGameStore } from "@/stores/game.store"
import { EWSAction } from "./actions"
import { EWSMessage, TUpdateGameStateMsg } from "./messages"
import { TWsService, wsService } from "./ws.service"

class GameService {
  private readonly ws: TWsService
  constructor() {
    this.ws = wsService
    this.ws.listen(EWSMessage.UpdateGameState, this.onUpdateGameState)
  }

  public startGame = () => {
    this.ws.send({ type: EWSAction.StartGame })
  }

  private onUpdateGameState = (data: TUpdateGameStateMsg) => {
    const gameStore = useGameStore.getState()
    gameStore.updateGame(data.game)
  }
}

export const gameService = new GameService()