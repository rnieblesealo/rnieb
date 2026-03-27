import { useEffect, useState } from "react"
import AuthContext from "./AuthContext"
import axios from "axios"

export default function AuthContextProvider({ children }: { children: React.ReactNode }) {
  const [loggedIn, setLoggedIn] = useState(false)

  useEffect(() => {
    // check auth and set login state every remount
    axios.get(`${import.meta.env.VITE_BASE_URL}/me`)
      .then(() => setLoggedIn(true))
      .catch(() => setLoggedIn(false))
  }, [])

  return (
    // here we give the context the actual values
    <AuthContext.Provider value={{
      loggedIn,
      setLoggedIn
    }}>
      {children}
    </AuthContext.Provider>
  )
}

