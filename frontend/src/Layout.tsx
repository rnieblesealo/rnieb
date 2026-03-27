import axios from "axios"
import { useEffect, useState } from "react"
import { Outlet } from "react-router-dom"
import NavButton from "./components/NavButton"
import useAuthContext from "./context/useAuthContext"

export default function Layout() {
  const { loggedIn, setLoggedIn } = useAuthContext()

  const [logo, setLogo] = useState('')
  const [deco, setDeco] = useState('')

  async function handleLogin(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault()

    // submit login credentials

    const formData = new FormData(e.currentTarget)
    axios.post(`${import.meta.env.VITE_BASE_URL}/login`, formData)
      .then(res => {
        // save login token to localstorage; set in axios defaults
        // see the tradeoffs of this in README

        const token = res.data.data.token
        localStorage.setItem("token", token)

        axios.defaults.headers.common["Authorization"] = token

        setLoggedIn(true)
      })
  }

  function handleLogout() {
    // remove token from local storage

    localStorage.removeItem("token")

    // delete axios auth header
    /* delete keyword removes a property from an object */

    delete axios.defaults.headers.common["Authorization"]

    setLoggedIn(false)
  }

  useEffect(() => {
    // get logo parts
    axios.get("/logo.txt").then(res => { setLogo(res.data) })
    axios.get("/stars.txt").then(res => { setDeco(res.data) })
  }, [])

  return (
    <div className="relative w-full h-min flex flex-col items-center justify-center">
      {/* fire skull bg */}
      <img
        src="frutiger-metro.png"
        className="absolute w-128 h-auto top-0 right-0 opacity-60 z-0"
      />

      {/* ascii title */}
      <div className="flex flex-row gap-5 text-xs z-99 mt-6">
        <pre>{deco}</pre>
        <pre>{logo}</pre>
      </div>

      {/* navbar */}
      <nav className="grid grid-cols-3 sm:grid-cols-4 gap-1 mb-8 mt-4 w-fit mx-auto">
        <NavButton route="/art" iconSrc="/icons/btn-art.png" />
        <NavButton route="/skate" iconSrc="/icons/btn-skate.png" />
        <NavButton route="/music" iconSrc="/icons/btn-music.png" />
        <NavButton route="/programming" iconSrc="/icons/btn-programming.png" />
      </nav>

      { /* page content */}
      <Outlet />

      {/* auth status */}
      <div className="mt-4 z-999">
        {!loggedIn ?
          // menu when NOT LOGGED 
          <div className="flex flex-col items-center justify-center">
            <form
              onSubmit={handleLogin}
              className="flex flex-col items-center justify-center text-center"
            >
              <span className="text-red-500 p-2">Not Logged In</span>
              <input
                type="text"
                name="username"
                placeholder="Username"
                className="w-full border border-red-900 hover:border-red-700 focus:border-red-500 px-2 py-1 my-1"
              />
              <input
                type="password"
                name="password"
                placeholder="Password"
                className="w-full border border-red-900 hover:border-red-700 focus:border-red-500 px-2 py-1 my-1"
              />
              <button
                type="submit"
                className="w-fit p-2 text-red-700 hover:text-red-600 active:text-red-500 text-shadow-red-600"
              >
                [ Log In ]
              </button>
            </form>
          </div> :
          // menu when LOGGED 
          <div className="flex flex-col items-center justify-center">
            <span className="text-green-500 font-bold p-2">Logged In!</span>
            <button
              onClick={handleLogout}
              className="w-fit p-2 text-red-700 hover:text-red-600 active:text-red-500 text-shadow-red-600"
            >
              [ Log Out ]
            </button>
          </div>
        }
      </div>

      {/* footer */}
      <footer className="w-fit h-fit">
        <img src="chiikawa.png" className="w-16 m-8" />
      </footer>

    </div >
  )
}
