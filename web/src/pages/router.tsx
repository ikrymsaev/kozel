import { createBrowserRouter } from "react-router-dom";
import { AuthPage } from "./AuthPage";
import { GamePage } from "./GamePage";
import { LobbyPage } from "./LobbyPage";

export const router = createBrowserRouter([
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