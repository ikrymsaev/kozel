import { useState } from "react"
import { useChatStore } from "../../../stores/chat.store"
import { lobbyService } from "../../../services/lobby.service"

export const LobbyChat = () => {
  const messages = useChatStore((state) => state.messages)

  const [input, setInput] = useState("")

  return (
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
      
      <div className="flex flex-col gap-1">
        {messages.map((msg, i) => (
          <p key={i} className="text-sm px-1">
            {msg}
          </p>
        ))}
      </div>
    </div>
  )
}