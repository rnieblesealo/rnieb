import { useState } from "react"
import type { Drawing } from "../types"
import axios from "axios"

interface DrawingTileProps {
  data: Drawing,
  loggedIn: boolean,
  onDelete: () => void
}

const DrawingTile = ({ data, loggedIn, onDelete }: DrawingTileProps) => {
  const [hovered, setHovered] = useState(false)

  return (
    // entire tile 
    <div
      className="w-48 flex flex-col items-center justify-center"
      onMouseEnter={() => { setHovered(true) }}
      onMouseLeave={() => { setHovered(false) }}
    >
      <div className="relative">
        {/* image itself */}
        <img
          src={`${import.meta.env.VITE_BASE_URL}/uploads/${data.path}`}
          className="w-fit aspect-square object-cover"
        />
        {/* description overlay */}
        {hovered &&
          <div
            className="absolute top-0 opacity-80 w-full h-full text-white text-sm bg-black flex flex-col items-center justify-center">
            <span className="italic font-bold mt-4 mb-[-1] text-md text-center">"{data.name}"</span>
            <span className="opacity-75 m-4 text-center text-sm overflow-auto">
              {data.description}
            </span>
            {/* delete button */}
            {loggedIn &&
              <button
                onClick={() => {
                  axios.delete(`${import.meta.env.VITE_BASE_URL}/delete-drawing`, {
                    params: { id: data.id }
                  }).then(() => onDelete())
                }}
                className="relative w-fit p-2 mb-4 bg-red-500">
                Delete
              </button>
            }
          </div>
        }
      </div>
    </div>
  )
}

export default DrawingTile
