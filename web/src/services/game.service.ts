import { EWSAction } from "./actions"
import { TWsService, wsService } from "./ws.service"

class GameService {
  private readonly ws: TWsService
  constructor() {
    this.ws = wsService
  }

  public startGame = () => {
    this.ws.send({ type: EWSAction.StartGame })
  }
}

export const gameService = new GameService()