import axios from "axios"
import { useEffect, useState } from "react"

interface GetDrawingsResponse {
  message: string,
  data: Drawing[]
}

interface Drawing {
  id: string,
  name: string,
  description: string,
  path: string
}

const DrawingTile = ({ data, onDelete }: { data: Drawing, onDelete: () => void }) => {
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
            <span className="italic font-bold mt-4 mb-[-1] text-md text-center">"{data.name}"</span>
            <span className="opacity-75 m-4 text-center text-sm overflow-auto">
              {data.description}
            </span>
            {/* delete button */}
            <button
              onClick={() => {
                axios.delete("http://localhost:8080/delete-drawing", {
                  params: { id: data.id }
                }).then(() => onDelete())
              }}
              className="relative w-fit p-2 mb-4 bg-red-500">
              Delete
            </button>
          </div>
        }
      </div>
    </div>
  )
}

interface CollageProps {
  drawings: Drawing[],
  setDrawings: React.Dispatch<React.SetStateAction<Drawing[]>>
}

const Collage = ({ drawings, setDrawings }: CollageProps) => {

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
          onDelete={() => handleDelete(drawing.id)}
        />
      ))}
    </div>
  )
}

// retrieve and set auth token

/* this stuff runs on app entry point before anything is rendered
useffects inside <app> run on component mount
for auth shit do this outside since we want it done before anything renders */

const token = localStorage.getItem("token")
if (token) {
  axios.defaults.headers.common["Authorization"] = token
}

export default function App() {
  const [logo, setLogo] = useState('')
  const [deco, setDeco] = useState('')
  const [drawings, setDrawings] = useState<Drawing[]>([])
  const [loggedIn, setLoggedIn] = useState(false)

  useEffect(() => {
    // get logo parts

    axios.get("/logo.txt").then(res => { setLogo(res.data) })
    axios.get("/stars.txt").then(res => { setDeco(res.data) })

    // fetch drawings

    axios.get("http://localhost:8080/get-drawings")
      .then(res => {
        const getDrawingsResponse: GetDrawingsResponse = res.data
        setDrawings(getDrawingsResponse.data)
      })
      .catch(err => {
        console.error("Failed to fetch images:", err)
      })

    // check auth

    // console.log(localStorage.getItem("token"))

    axios.get("http://localhost:8080/me")
      .then(() => setLoggedIn(true))
  }, [])


  // use custom submission function to avoid json response page behavior

  async function handleUploadForm(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault()

    // submit image upload request 

    const formData = new FormData(e.currentTarget)
    axios.post("http://localhost:8080/upload", formData) // submit upload
      .then(() => {
        // on upload complete, refetch drawings to update display

        axios.get("http://localhost:8080/get-drawings")
          .then(res => {
            const getDrawingsResponse: GetDrawingsResponse = res.data
            setDrawings(getDrawingsResponse.data)
          })
      })
  }

  async function handleLogin(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault()

    // submit login credentials

    const formData = new FormData(e.currentTarget)
    axios.post("http://localhost:8080/login", formData)
      .then(res => {
        // save login token to localstorage
        // see the tradeoffs of this in README

        const token = res.data.data.token
        localStorage.setItem("token", token)
      })
  }

  return (
    <div className="w-full h-min flex flex-col items-center justify-center">
      {/* auth status */}
      <span className="text-red-500 m-1">Not Logged In</span>
      <form
        onSubmit={handleLogin}
        className="flex flex-col items-start gap-3 m-1"
      >
        <input
          type="text"
          name="username"
          placeholder="Username"
          className="w-full"
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          className="w-full"
        />
        <button
          type="submit"
          className="bg-red-500 text-black p-1 w-full font-bold "
        >
          Log In
        </button>
      </form>

      {/* ascii title */}
      <div className="flex flex-row gap-5 text-xs m-4">
        <pre>{deco}</pre>
        <pre>{logo}</pre>
      </div>

      {/* display images */}
      <div>
        <Collage
          drawings={drawings}
          setDrawings={setDrawings}
        />
      </div >

      {/* upload image form */}
      <div className="flex flex-col items-center justify-center m-4 text-red-500">
        <span className="mb-4">Upload a Drawing</span>
        <form
          onSubmit={handleUploadForm}
          className="flex flex-col items-start gap-3 m-4"
        >
          {/* image uploader */}
          <input
            id="upload-images"
            type="file"
            name="file"
            accept="image/*"
            className="w-full"
          />
          {/* image name */}
          <input
            type="text"
            name="name"
            placeholder="Name..."
            className="w-full"
          />
          {/* image description */}
          <textarea
            name="description"
            placeholder="Description..."
            rows={4}
            className="w-full"
          />
          {/* submit button */}
          <button
            type="submit"
            className="bg-red-500 text-black p-2 w-full font-bold "
          >
            Upload
          </button>
        </form>
      </div>
    </div>
  )
}
