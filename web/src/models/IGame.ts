import { IPlayer } from "./IPlayer";
import { IRound } from "./IRound";

export interface IGameState {
  players: [IPlayer, IPlayer, IPlayer, IPlayer];
  round: IRound;
  score: [number, number];
}