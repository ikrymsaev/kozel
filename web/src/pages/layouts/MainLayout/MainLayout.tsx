import { PropsWithChildren } from "react"
import styles from './MainLayout.module.css'
import cn from 'classnames'

export const MainLayout = ({ children }: PropsWithChildren) => {

  return (
    <div className={cn(styles.container, "p-4")}>
      <div>
        <h5>Kozel (MainLayout)</h5>
      </div>
      <div>
        {children}
      </div>
    </div>
  )
}