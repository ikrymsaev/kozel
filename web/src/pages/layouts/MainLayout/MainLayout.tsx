import { PropsWithChildren } from "react"
import styles from './MainLayout.module.css'

export const MainLayout = ({ children }: PropsWithChildren) => {

  return (
    <div className={styles.container}>
      <div>
        <h1>Kozel (MainLayout)</h1>
      </div>
      <div>
        {children}
      </div>
    </div>
  )
}