import type { Drawing } from "../types"
import DrawingTile from "./DrawingTile"

interface CollageProps {
  drawings: Drawing[],
  loggedIn: boolean,
  setDrawings: React.Dispatch<React.SetStateAction<Drawing[]>>
}

const Collage = ({ drawings, loggedIn, setDrawings }: CollageProps) => {

  /* when a tile is deleted, we filter it out of the list to reflect deletion
   * it will be gone on refresh fully since the useffect will fire */

  const handleDelete = (id: string) => {
    setDrawings(prev => prev.filter(d => d.id !== id))
  }

  return (
    /* render the images */
    <div className="grid grid-cols-2 md:grid-cols-3 w-fit">
      {drawings && drawings.map(drawing => (
        <DrawingTile
          key={drawing.id}
          data={drawing}
          loggedIn={loggedIn}
          onDelete={() => handleDelete(drawing.id)}
        />
      ))}
    </div>
  )
}

export default Collage
