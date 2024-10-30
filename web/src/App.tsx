import './App.css'
import { RouterProvider } from 'react-router-dom'
import { router } from './pages/router'
import { Bounce, ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css';

function App() {
  return (
    <>
      <RouterProvider router={router} />
      <ToastContainer
        position="bottom-right"
        closeButton={false}
        autoClose={1000}
        hideProgressBar
        closeOnClick
        rtl={false}
        pauseOnFocusLoss={false}
        draggable
        pauseOnHover={false}
        theme="dark"
        transition={Bounce}
      />
    </>
  )
}

export default App
