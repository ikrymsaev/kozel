import { LogoutIcon } from "@/shared/icons/LogoutIcon"
import { PersonIcon } from "@/shared/icons/PersonIcon"
import { Text } from "@/shared/ui-kit/Text"
import { useAuthStore } from "@/stores/auth.store"
import { useNavigate } from "react-router-dom"

export const Header = () => {
  const user = useAuthStore((state) => state.user)
  const signOut = useAuthStore((state) => state.signOut)

  const navigate = useNavigate()

  if (!user) return null

  return (
    <div className="flex items-center justify-between w-full px-4 py-2 bg-emerald-900 min-h-10">
      <div className="flex flex-1 justify-between">
        <div
          className="flex items-center gap-2 cursor-pointer active:scale-95"
          onClick={() => navigate("/")}
        >
          <PersonIcon />
          <Text className="italic font-serif text-nowrap">
            {user.username}
          </Text>
        </div>
        <LogoutIcon
          onClick={signOut}
          className="cursor-pointer active:scale-95"
        />
      </div>
    </div>
  )
}