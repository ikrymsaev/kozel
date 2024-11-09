import { createHashRouter } from "react-router-dom";
import { AuthProvider } from "./components/providers/AuthProvider";
import { AuthPage } from "./components/pages/AuthPage";
import { MainPage } from "./components/pages/MainPage";
import { LobbyPage } from "./components/pages/LobbyPage/LobbyPage";
import { GamePage } from "./components/pages/GamePage";

export const router = createHashRouter([
  {
    path: "/auth",
    element: <AuthPage />,
  },
  {
    path: "/",
    element: <AuthProvider><MainPage /></AuthProvider>,
  },
  {
    path: "/lobby",
    element: <AuthProvider><LobbyPage /></AuthProvider>,
  },
  {
    path: "/game",
    element: <AuthProvider><GamePage /></AuthProvider>,
  },
]);