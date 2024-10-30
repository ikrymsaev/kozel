import { ILobbySlot } from "../../../models/ILobby"
import { useAuthStore } from "../../../stores/auth.store"
import { useLobbyStore } from "../../../stores/lobby.store"
import cn from "classnames"

export const LobbySlots = () => {
  const slots = useLobbyStore((state) => state.slots)

  return (
    <div>
      <h5>Lobby Players</h5>
      <div className="flex flex-col gap-2 border-2 rounded-sm">
        {slots.map((slot) => <Slot key={slot.order} slot={slot} />)}
      </div>
    </div>
  )
}

const Slot = ({ slot }: { slot: ILobbySlot }) => {
  const user = useAuthStore((state) => state.user)
  const isMe = slot.player?.id === user?.id

  if (!slot.player) {
    return (
      <div className="px-4 py-2 rounded-sm bg-slate-300 text-gray-600">
        Free slot
      </div>
    )
  }

  return (
    <div
      className={cn(
        isMe && 'font-semibold italic border-[1px] border-slate-800',
        (slot.order % 2) ? "bg-sky-400" : "bg-red-400",
        "px-4 py-2 rounded-md text-slate-900",
      )}
    >
      {slot.player.username}
    </div>
  )
}