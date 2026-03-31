// This is a more basic layout that doesnt have navbar for 404 page to cleanly display game
// Maybe rename to something more broad if you end up using it for other stuff

import { useEffect, useState } from "react"
import axios from "axios"
import { Outlet } from "react-router-dom"

export default function NotFoundLayout() {
  const [logo, setLogo] = useState('')
  const [deco, setDeco] = useState('')

  useEffect(() => {
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

      { /* page content */}
      <Outlet />
    </div >
  )
}
