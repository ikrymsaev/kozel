import { HTMLAttributes, PropsWithChildren } from "react"
import styles from './MainLayout.module.css'
import cn from 'classnames'

export const MainLayout = ({ children, ...props }: PropsWithChildren & HTMLAttributes<HTMLDivElement>) => {

  return (
    <div {...props} className={cn(props.className, styles.container, "p-4 w-full max-w-md mx-auto h-screen overflow-hidden bg-green-800 text-white")}>
      {children}
    </div>
  )
}