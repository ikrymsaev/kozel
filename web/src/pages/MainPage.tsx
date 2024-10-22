import { useEffect, useState } from "react";
import { lobbyService } from "../services/lobby.service";
import { useLobbyStore } from "../stores/lobby.store";
import { useNavigate } from "react-router-dom";
import { ILobby } from "../models/ILobby";

export const MainPage = () => {
  const navigate = useNavigate()

  const activeGames = useLobbyStore((state) => state.activeGames)
  const addLobby = useLobbyStore((state) => state.addLobby)
  
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    lobbyService.getLobbyList()
  }, [])

  useEffect(() => {
    const es = new EventSource('http://localhost:8080/hub/watch_lobbies')
    console.log("SSE opened, watching for lobbies...")
    es.addEventListener('new_lobby', (event) => {
      const data = JSON.parse(event.data) as ILobby
      console.log("SSE new_lobby: ", data)
      addLobby(data)
    })

    return () => {
      es.close()
    }
  }, [addLobby])

  const handleNewGame = async () => {
    try {
      setLoading(true);
      await lobbyService.newGame();
    }
    finally {
      setLoading(false);
    }
  }

  const enterGame = async (gameId: string) => {
    navigate('/lobby?id=' + gameId)
  }

  return (
    <div>
      <h1>Main Page</h1>
      <button disabled={loading} onClick={handleNewGame}>new game</button>
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