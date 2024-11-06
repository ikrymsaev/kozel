import { IPlayer } from "./IPlayer";
import { IRound } from "./IRound";

export enum EGameStage {
  Preparing,
  Praising,
  PlayerStep,
  Calculation,
  DealerStep,
  GameOver
}

export interface IGameState {
  round: IRound;
  stage: EGameStage;
  score: [number, number];
  players: [IPlayer, IPlayer, IPlayer, IPlayer];
}
