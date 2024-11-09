import { lobbyService } from "@/services/lobby.service"
import { getApiUrl } from "@/shared/utils/get-api-url"
import { useAuthStore } from "@/stores/auth.store"
import { useLobbyStore } from "@/stores/lobby.store"
import { useEffect } from "react"

export const useConnectToHub = () => {
  const token = useAuthStore((state) => state.token)
  const removeLobby = useLobbyStore((state) => state.removeLobby)

  useEffect(() => {
    if (!token) return
    const ws = new WebSocket(getApiUrl('ws') + "/api/hub/connect/" + token)
    ws.onmessage = (e: { data: string }) => {
      const message = JSON.parse(e.data) as Record<string, any>
      console.log('HUB: ', message)
      if (message?.type === 0) { // new lobby
        lobbyService.fetchNewLobby(message.id)
      }
      else if (message?.type === 1) { // remove lobby
        removeLobby(message.id)
      }
    }

    return () => ws.close()
  }, [removeLobby, token])
}