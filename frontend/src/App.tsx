import axios from "axios"
import { useEffect, useState } from "react"

interface Drawing {
  id: string,
  name: string,
  description: string,
  path: string
}

const DrawingTile = ({ data }: { data: Drawing }) => {
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
          src={`http://localhost:8080/${data.path}`}
          className="w-fit aspect-square object-cover"
        />
        {/* description overlay */}
        {hovered &&
          <div
            className="absolute top-0 opacity-80 w-full h-full text-white text-sm bg-black flex flex-col items-center justify-center">
            <span className="italic font-bold mt-4 mb-[-1] text-md">"{data.name}"</span>
            <span className="opacity-75 m-4 text-center text-sm overflow-auto">
              {data.description}
            </span>
          </div>
        }
      </div>
    </div>
  )
}

const Collage = () => {
  const [drawings, setDrawings] = useState<Drawing[]>([])

  // fetch images
  useEffect(() => {
    axios.get("http://localhost:8080/get-drawings")
      .then(res => {
        setDrawings(res.data)
      })
      .catch(err => {
        console.error("Failed to fetch images:", err)
      })
  }, [])

  return (
    /* render the images */
    <div className="grid grid-cols-2 md:grid-cols-3 w-fit">
      {drawings && drawings.map(drawing => (
        <DrawingTile key={drawing.id} data={drawing} />
      ))}
    </div>
  )
}

export default function App() {
  const [logo, setLogo] = useState('')
  const [deco, setDeco] = useState('')

  axios.get("/logo.txt").then(res => { setLogo(res.data) })
  axios.get("/stars.txt").then(res => { setDeco(res.data) })

  return (
    <div className="w-full h-min flex flex-col items-center justify-center">
      <div className="flex flex-row gap-5 text-xs m-6">
        <pre>{deco}</pre>
        <pre>{logo}</pre>
      </div>

      {/* display images */}
      <div>
        <Collage />
      </div >

      {/* upload image form */}
      <div className="flex flex-col items-center justify-center m-6 text-red-500">
        <span className="font-bold mb-6">Upload a Drawing</span>
        <form
          action="http://localhost:8080/upload"
          method="POST"
          encType="multipart/form-data"
          className="flex flex-col items-start gap-3 m-4"
        >
          <input
            type="text"
            name="name"
            placeholder="Name..."
            className="w-full"
          />
          <textarea
            name="description"
            placeholder="Description..."
            rows={4}
            className="w-full"
          />
          <input
            id="upload-images"
            type="file"
            name="file"
            accept="image/*"
            className="w-full"
          />
          <button type="submit" className="bg-red-500 text-black p-2 w-full font-bold ">Upload</button>
        </form>
      </div>
    </div>
  )
}
