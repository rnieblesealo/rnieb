import { createContext } from "react"

interface AuthContextValues {
  loggedIn: boolean,
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

// these are the default values of the context
const AuthContext = createContext<AuthContextValues>({
  loggedIn: false,
  setLoggedIn: () => { }
})

export default AuthContext
