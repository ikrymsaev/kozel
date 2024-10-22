import { useEffect, useState } from "react";
import { lobbyService } from "../services/lobby.service";
import { useLobbyStore } from "../stores/lobby.store";
import { useNavigate } from "react-router-dom";

export const MainPage = () => {
  const navigate = useNavigate()

  const activeGames = useLobbyStore((state) => state.activeGames)
  
  const [creating, setCreating] = useState(false);

  useEffect(() => {
    const unwatch = lobbyService.watchLobbies()
    return unwatch
  }, [])

  const handleNewGame = async () => {
    try {
      setCreating(true);
      await lobbyService.newGame();
    }
    finally {
      setCreating(false);
    }
  }

  const enterGame = async (gameId: string) => {
    navigate('/lobby?id=' + gameId)
  }

  return (
    <div>
      <h1>Main Page</h1>
      <button disabled={creating} onClick={handleNewGame}>new game</button>
      <div>
        {activeGames.map(game => (
          <div key={game.id} onClick={() => enterGame(game.id)} style={{ cursor: 'pointer' }}>
            {game.id} | {game.name}
          </div>
        ))}
      </div>
    </div>
  )
}