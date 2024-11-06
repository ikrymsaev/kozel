import { createHashRouter } from "react-router-dom";
import { AuthPage } from "./AuthPage";
import { GamePage } from "./GamePage";
import { LobbyPage } from "./LobbyPage/LobbyPage";
import { MainPage } from "./MainPage";

export const router = createHashRouter([
  {
    path: "/",
    element: <MainPage />,
  },
  {
    path: "/auth",
    element: <AuthPage />,
  },
  {
    path: "/lobby",
    element: <LobbyPage />,
  },
  {
    path: "/game",
    element: <GamePage />,
  },
]);