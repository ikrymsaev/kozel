import { HTMLAttributes, PropsWithChildren } from "react"
import styles from './MainLayout.module.css'
import cn from 'classnames'
import { Header } from "@/components/widgets/Header"

export const MainLayout = ({ children, ...props }: PropsWithChildren & HTMLAttributes<HTMLDivElement>) => {

  return (
    <div {...props} className={cn(props.className, styles.container, "w-full max-w-md mx-auto h-screen overflow-hidden bg-green-800 text-white")}>
      <Header />
      {children}
    </div>
  )
}