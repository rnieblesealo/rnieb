import './index.css'
import Art from './pages/Art'
import AuthContextProvider from './context/AuthContextProvider'
import Layout from './Layout'
import axios from 'axios'
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom"
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

// pull and set auth token every reload
// want this done before anything renders

const token = localStorage.getItem("token")
if (token) {
  axios.defaults.headers.common["Authorization"] = token
}

// multiple <Routes> in same place = different sections ( e.g. login vs actual app )

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthContextProvider> {/* allows calling useAuthContext to set login value globally */}

      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />} >

            {/* default route is /art; we get redirected to it upon hitting / */}

            {/* replace = remove / history entry such that back button 
            doesn't take us to / again */}

            <Route index element={<Navigate to="/art" replace />} />
            <Route path="art" element={<Art />} />
          </Route>
        </Routes>
      </BrowserRouter>

    </AuthContextProvider>
  </StrictMode>,
)
