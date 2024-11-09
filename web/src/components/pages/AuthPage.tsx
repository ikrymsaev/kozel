import { Button } from "@/shared/ui-kit/Button"
import { Text } from "@/shared/ui-kit/Text"
import { useAuthStore } from "@/stores/auth.store"
import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { toast } from "react-toastify"

export const AuthPage = () => {
  const navigate = useNavigate()

  const signUp = useAuthStore((state) => state.signUp)
  const signIn = useAuthStore((state) => state.signIn)
  const loading = useAuthStore((state) => state.loading)

  const user = useAuthStore((state) => state.user)
  const token = useAuthStore((state) => state.token)

  const [formType, setFormType] = useState<"signUp" | "signIn">("signIn")

  useEffect(() => {
    if (user && token) {
      navigate('/')
    }
  }, [user, token, navigate])

  const handleSignUp = async (e: any) => {
    e.preventDefault()
    console.log(e)
    const formData = new FormData(e.target)
    const username = formData.get('username')?.toString()
    const password = formData.get('password')?.toString()
    if (!username || !password) {
      toast("incorrect username or password", { type: "error" })
      return;
    }
    const dto = { username, password }
    console.log(dto)
    await signUp(dto)
    navigate('/')
  }

  const handleSignIn = async (e: any) => {
    e.preventDefault()
    console.log(e)
    const formData = new FormData(e.target)
    const username = formData.get('username')?.toString()
    const password = formData.get('password')?.toString()
    if (!username || !password) {
      toast("incorrect username or password", { type: "error" })
      return;
    }
    const dto = { username, password }
    console.log(dto)
    await signIn(dto)
    navigate('/')
  }

  return (
    <div className="flex flex-col flex-grow justify-center w-full p-4 bg-emerald-900 text-white min-h-screen">
      <Text type="header" className="text-center">
        {formType === "signUp" ? "Регистрация" : "Войти"}
      </Text>
      <div className="flex flex-col gap-4 p-4">
        {formType === "signIn" && (
          <form onSubmit={handleSignIn} className="flex flex-col gap-2 items-center w-full flex-1">
            <input type="text" name="username" id="username" />
            <input type="password" name="password" id="password" />
            <Button type="submit" size="m" color="gold" loading={loading}>Sign In</Button>
          </form>
        )}
        {formType === "signUp" && (
          <form onSubmit={handleSignUp} className="flex flex-col gap-2 items-center w-full">
            <input type="text" name="username" id="username" />
            <input type="password" name="password" id="password" />
            <Button type="submit" size="m" color="gold" loading={loading}>Sign Up</Button>
          </form>
        )}
      </div>
      <div className="flex justify-center mt-10">
        <Button
          mode="bezeled"
          color="transparent"
          loading={loading}
          onClick={() => setFormType(formType === "signUp" ? "signIn" : "signUp")}
        >
          {formType === "signUp" ? "войти" : "регистрация"}
        </Button>
      </div>
    </div>
  )
}