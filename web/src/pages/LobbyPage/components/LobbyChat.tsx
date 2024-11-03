import { useEffect, useRef, useState } from "react"
import { chatService } from "@/services/chat.service"
import { useChatStore } from "@/stores/chat.store"
import { IChatMessage } from "@/models/IChatMessage"
import { useAuthStore } from "@/stores/auth.store"
import cn from "classnames"

export const LobbyChat = () => {
  const me = useAuthStore((state) => state.user)
  const messages = useChatStore((state) => state.messages)

  const [input, setInput] = useState("")

  const chatRef = useRef<HTMLDivElement>(null)

  const isMineMessage = (msg: IChatMessage) => {
    return msg.sender.id === me?.id
  }

  useEffect(() => {
    const observer = new MutationObserver((mutationsList: MutationRecord[]) => {
      for (const mutation of mutationsList) {
        if (mutation.type === 'childList') {
          return chatRef.current!.scrollTo(0, chatRef.current!.scrollHeight)
        }
      }
    })
    observer.observe(chatRef.current!, { childList: true })
  }, [])

  return (
    <div className="flex flex-col justify-end flex-grow gap-2">
      <div
        ref={chatRef}
        className="flex flex-col flex-grow gap-1 min-h-28 max-h-48 max-w-sm overflow-auto bg-slate-300 rounded-sm"
      >
        {messages.map((msg, i) => (
          <div key={`msg-${i}`}>
            <span
              style={{ color: getSenderColor(msg) }}
              className={cn(
                "text-sm italic font-semibold"
              )}
            >
              {isMineMessage(msg) ? "" : `${msg.sender.username}: `}
            </span>
            <span
              className={cn(
                "text-slate-800",
                isMineMessage(msg) ? "text-right" : "text-left"
              )}>
              {msg.message}
            </span>
          </div>
        ))}
      </div>
      
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
            chatService.sendChatMessage(input)
            setInput("")
          }}
        >
          {">"}
        </button>
      </form>
    </div>
  )
}

const colors: Record<string, string> = {}

const getSenderColor = (msg: IChatMessage) => {
  if (colors[msg.sender.id]) return colors[msg.sender.id]

  const color = Math.floor(Math.random() * 16777215).toString(16)
  colors[msg.sender.id] = color
  return color
}