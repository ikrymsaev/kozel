import { useChatStore } from "@/stores/chat.store"
import { EWSMessage, TNewMessageMsg } from "../api/ws/ws.messages"
import { TWsService, wsService } from "../api/ws/ws.service"
import { useAuthStore } from "@/stores/auth.store"
import { EWSAction, TSendMessage } from "../api/ws/ws.actions"

class ChatService {
  private readonly ws: TWsService
  constructor() {
    this.ws = wsService
    this.ws.listen(EWSMessage.NewMessage, this.onNewMessage)
  }

  /** Send message to chat */
  public sendChatMessage = (message: string) => {
    const user = useAuthStore.getState().user
    if (!user) return
    this.ws.send<TSendMessage>({
      type: EWSAction.SendMessage,
      message
    })
  }

  private onNewMessage = (m: TNewMessageMsg) => {
    const chatStore = useChatStore.getState()
    chatStore.addMessage(m)
  }
}

export const chatService = new ChatService()
