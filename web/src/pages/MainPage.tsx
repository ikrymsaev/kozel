import { useEffect, useState } from "react";
import { lobbyService } from "../services/lobby.service";
import { useLobbyStore } from "../stores/lobby.store";
import { useNavigate } from "react-router-dom";
import { Button } from "@/shared/ui-kit/Button";
import { Text } from "@/shared/ui-kit/Text";
import { MainLayout } from "./layouts/MainLayout/MainLayout";

export const MainPage = () => {
  const navigate = useNavigate()

  const activeGames = useLobbyStore((state) => state.activeGames)
  
  const [creating, setCreating] = useState(false);

  useEffect(() => {
    const unwatch = lobbyService.watchLobbies()
    return () => {
      unwatch()
    }
  }, [])

  const handleNewGame = async () => {
    try {
      setCreating(true);
      const lobbyId = await lobbyService.newLobby();
      if (!lobbyId) return
      navigate('/lobby?id=' + lobbyId)
    }
    finally {
      setCreating(false);
    }
  }

  const enterGame = (gameId: string) => {
    navigate('/lobby?id=' + gameId)
  }

  return (
    <MainLayout>
      <div className="flex flex-col items-center gap-4 w-full overflow-hidden">
        <Text type="header" className="font-semibold">Main Page</Text>
        <Button size="m" disabled={creating} onClick={handleNewGame}>
          NEW GAME
        </Button>
        <div className="flex flex-col gap-1 w-full">
          {activeGames.map(lobby => (
            <div
              key={lobby.id}
              onClick={() => enterGame(lobby.id)} style={{ cursor: 'pointer' }}
              className="text-sky-700 px-4 py-2 rounded-md bg-slate-300 hover:bg-slate-400"
            >
              <Text>
                {lobby.name}
              </Text>
            </div>
          ))}
        </div>
      </div>
    </MainLayout>
  )
}