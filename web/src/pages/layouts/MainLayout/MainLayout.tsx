import { PropsWithChildren } from "react"
import styles from './MainLayout.module.css'

export const MainLayout = ({ children }: PropsWithChildren) => {

  return (
    <div className={styles.container}>
      <div>
        <h5>Kozel (MainLayout)</h5>
      </div>
      <div>
        {children}
      </div>
    </div>
  )
}