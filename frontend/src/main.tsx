import './index.css'

import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom"
import axios from 'axios'

import AuthContextProvider from './context/AuthContextProvider'
import Layout from './layout/Layout'
import Art from './pages/Art'
import Skate from './pages/Skate'
import Music from './pages/Music'
import Programming from './pages/Programming'
import Niko from './pages/Niko'
import NotFoundLayout from './layout/NotFoundLayout'

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
            <Route path="/art" element={<Art />} />

            <Route path="/skate" element={<Skate />} />
            <Route path="/music" element={<Music />} />
            <Route path="/programming" element={<Programming />} />

          </Route>
          {/* 404 displays niko game */}
          <Route path="*" element={<NotFoundLayout />}>
            <Route path="*" element={<Niko />} />
          </Route>

        </Routes>
      </BrowserRouter>

    </AuthContextProvider>
  </StrictMode>,
)
