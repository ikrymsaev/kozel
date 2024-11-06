import { ESuit, ICard } from "@/models/ICard";
import { EGameStage, IGameState } from "@/models/IGame";
import { create } from "zustand";
import { immer } from "zustand/middleware/immer";


interface State {
  game: IGameState | null
  table: ICard[]
  isGameOver: boolean
}
interface Actions {
  gameOver: () => void;
  setRoundResult: (winnerTeam: number, score: number) => void
  updateGame: (game: IGameState) => void
  changeTurn: (playerId: string) => void
  updateStage: (stage: EGameStage) => void
  setTrump: (trump: ESuit) => void
  reset: () => void
  getPlayerName: (playerId: string) => string
  replaceToTable: (card: ICard) => void
  moveFromTableToBribes: (winnerId: string, points: number) => void
}
type TGameStore = State & Actions;

export const useGameStore = create<TGameStore>()(immer((set, get) => ({
  /** State */
  game: null,
  table: [],
  isGameOver: false,
  /** Actions */
  gameOver: () => set((state) => {
    if (!state.game) return
    state.game.stage = EGameStage.GameOver
    state.isGameOver = true
  }),
  setRoundResult: (winnerTeam: number, score: number) => set((state) => {
    if (!state.game) return
    if (winnerTeam < 1 || winnerTeam > 2) {
      console.error("Invalid winnerTeam", winnerTeam)
      return
    }
    state.game.score[winnerTeam - 1] += score
  }),
  moveFromTableToBribes: (winnerId: string, points: number) => set((state) => {
    if (!state.game?.round) return
    const winner = state.game.players.find((player) => player.id === winnerId)
    if (!winner) return
    if (winner.position === 1 || winner.position === 3) {
      state.game.round.bribes[0] += points
    } else {
      state.game.round.bribes[1] += points
    }
    state.table = []
  }),
  replaceToTable: (card: ICard) => set((state) => {
    if (!state.game) return
    for (const player of state.game.players) {
      player.hand = player.hand.filter((c) => c.id !== card.id)
    }
    state.table.push(card)
  }),
  updateGame: (game: IGameState) => set((state) => {
    state.game = game
  }),
  changeTurn: (playerId: string) => set((state) => {
    if (!state.game?.round) return
    state.game.round.turnPlayerId = playerId
  }),
  setTrump: (trump: ESuit) => set((state) => {
    if (!state.game?.round) return
    state.game.round.trump = trump
    for (const player of state.game.players) {
      for (const card of player.hand) {
        if (!card || card.isHidden) continue
        if (card.suit === trump) {
          card.isTrump = true
        }
      }
    }
  }),
  updateStage: (stage: EGameStage) => set((state) => {
    console.log("updateStage", stage)
    if (!state.game) return
    state.game.stage = stage
  }),
  reset: () => set(() => ({
    game: null,
    table: []
  })),
  getPlayerName: (playerId: string): string => {
    const game = get().game
    if (!game) return ''

    const player = game.players.find((player) => player.id === playerId)
    if (!player) return ''

    const botPostfix = !player.user ? ' (BOT)' : ''

    return player.name + botPostfix
  }
})))