import { ILobbySlot } from "../../../models/ILobby"
import { useAuthStore } from "../../../stores/auth.store"
import { useLobbyStore } from "../../../stores/lobby.store"
import cn from "classnames"
import UpIcon from "/icons/up.svg"
import DownIcon from "/icons/down.svg"
import { lobbyService } from "../../../services/lobby.service"
import { useSearchParams } from "react-router-dom"

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
  const [searchParams] = useSearchParams();
  const lobbyId = searchParams.get("id");
  const user = useAuthStore((state) => state.user)
  const isMe = slot.player?.id === user?.id
  const isOwner = user?.id === lobbyId
  

  const canMoveUp = slot.order > 1
  const canMoveDown = slot.order < 4

  const moveUp = () => {
    lobbyService.moveSlot(slot.order, slot.order - 1)
  }

  const moveDown = () => {
    lobbyService.moveSlot(slot.order, slot.order + 1)
  }

  if (!slot.player) {
    return (
      <div
        className={cn(
          (slot.order % 2) ? "text-sky-700" : "text-red-700",
          "px-4 py-2 rounded-sm bg-slate-300"
        )}
      >
        Free slot
      </div>
    )
  }

  return (
    <div
      className={cn(
        isMe && 'font-semibold italic border-[1px] border-slate-800',
        (slot.order % 2) ? "bg-sky-400" : "bg-red-400",
        "flex flex-row flex-nowrap items-center px-4 py-2 rounded-md text-slate-900",
      )}
    >
      <div className="flex flex-1">
        <span>
          {slot.player.username}
        </span>
      </div>
      {isOwner && (
        <div className="flex flex-col items-center gap-1">
          <div
            onClick={moveUp}
            className={cn(
              !canMoveUp && "hover:scale-100 hover:border-slate-700 opacity-50 pointer-events-none",
              "flex flex-col items-center cursor-pointer bg-slate-300 border rounded-full p-1 border-slate-700",
              "hover:scale-105 hover:border-slate-900"
            )}
          >
            <img src={UpIcon} width="16" height="16" />
          </div>
          <div
            onClick={moveDown}
            className={cn(
              !canMoveDown && "hover:scale-100 hover:border-slate-700 opacity-50 pointer-events-none",
              "flex flex-col items-center cursor-pointer bg-slate-300 border rounded-full p-1 border-slate-700",
              "hover:scale-105 hover:border-slate-900"
            )}
          >
            <img src={DownIcon} width="16" height="16" />
          </div>
        </div>
      )}
    </div>
  )
}