import { MainLayout } from "./layouts/MainLayout/MainLayout"
import { useEffect, useState } from "react"
import { useSettingsStore } from "../stores/settings.store"
import { useSearchParams } from "react-router-dom"
import { lobbyService } from "../services/lobby.service"
import { useChatStore } from "../stores/chat.store"

export const LobbyPage = () => {
  const [searchParams] = useSearchParams();
  const lobbyId = searchParams.get("id");

  const fetchDeck = useSettingsStore((state) => state.fetchDeck)
  const messages = useChatStore((state) => state.messages)

  const [input, setInput] = useState("")

  useEffect(() => {
    fetchDeck()
  }, [fetchDeck])

  /** Connect to websocket chat */
  useEffect(() => {
    if (!lobbyId) return
    lobbyService.joinLobby(lobbyId)
    return lobbyService.leaveLobby
  }, [lobbyId])

  return (
    <MainLayout>
      <h1>Lobby Page</h1>
      <div>
        <form onSubmit={(e) => e.preventDefault()}>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
          />
          <button
            type="submit"
            onClick={(e) => {
              e.preventDefault()
              lobbyService.sendChatMessage(input)
              setInput("")
            }}
          >
            {">"}
          </button>
        </form>
      </div>
      <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start', paddingLeft: '10px' }}>
        {messages.map((msg, i) => <p key={i}>{msg}</p>)}
      </div>
    </MainLayout>
  )
}