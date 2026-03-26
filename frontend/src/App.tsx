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

interface DrawingTileProps {
  data: Drawing,
  loggedIn: boolean,
  onDelete: () => void
}

// get base url
// changes depending on whether this is prod or dev
const baseUrl = import.meta.env.VITE_BASE_URL
console.log("Using base URL:", baseUrl)

// retrieve and set auth token

/* this stuff runs on app entry point before anything is rendered
useffects inside <app> run on component mount
for auth shit do this outside since we want it done before anything renders */

const token = localStorage.getItem("token")
if (token) {
  axios.defaults.headers.common["Authorization"] = token
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
          src={`${baseUrl}/uploads/${data.path}`}
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
                  axios.delete(`${baseUrl}/delete-drawing`, {
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

interface NavButtonProps {
  iconSrc: string
}

const NavButton = ({ iconSrc }: NavButtonProps) => {
  const [hovered, setHovered] = useState(false)

  return (
    <button
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
      className="relative overflow-hidden rounded-sm"
    >
      <img
        src={iconSrc}
        className="w-28 h-fit"
        style={{
          filter: hovered ? undefined : "saturate(0)"
        }}
      />
      {/* gradient overlay on top for ios 6 look */}
      <div className="bg-gradient-to-b from-white to-black opacity-30 w-full h-full absolute top-0" />
    </button>
  )
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

    axios.get(`${baseUrl}/get-drawings`)
      .then(res => {
        const getDrawingsResponse: GetDrawingsResponse = res.data
        setDrawings(getDrawingsResponse.data)
      })
      .catch(err => {
        console.error("Failed to fetch images:", err)
      })

    // check auth

    axios.get(`${baseUrl}/me`)
      .then(() => setLoggedIn(true))
      .catch(() => setLoggedIn(false))
  }, [])


  // use custom submission function to avoid json response page behavior

  async function handleUploadForm(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault()

    // submit image upload request 

    const formData = new FormData(e.currentTarget)
    axios.post(`${baseUrl}/upload`, formData) // submit upload
      .then(() => {
        // on upload complete, refetch drawings to update display

        axios.get(`${baseUrl}/get-drawings`)
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
    axios.post(`${baseUrl}/login`, formData)
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
      <nav className="flex flex-row mb-8 mt-4 gap-1">
        <NavButton iconSrc="/icons/btn-art.png" />
        <NavButton iconSrc="/icons/btn-skate.png" />
      </nav>

      {/* display images */}
      <div>
        <Collage
          drawings={drawings}
          loggedIn={loggedIn}
          setDrawings={setDrawings}
        />
      </div>

      {/* upload form */}
      {
        loggedIn &&
        <div className="flex flex-col items-center justify-center text-red-500 mt-8">
          <span className="m-2">Upload a Drawing</span>
          <form
            onSubmit={handleUploadForm}
            className="flex flex-col items-center"
          >
            {/* image uploader */}
            <input
              id="upload-images"
              type="file"
              name="file"
              accept="image/*"
              className="w-full p-2"
            />
            {/* image name */}
            <input
              type="text"
              name="name"
              placeholder="Name..."
              className="w-full border border-red-900 hover:border-red-700 focus:border-red-500 px-2 py-1 my-1"
            />
            {/* image description */}
            <textarea
              name="description"
              placeholder="Description..."
              rows={4}
              className="w-full border border-red-900 hover:border-red-700 focus:border-red-500 px-2 py-1 my-1"
            />
            {/* submit button */}
            <button
              type="submit"
              className="w-fit p-2 text-red-700 hover:text-red-600 active:text-red-500 text-shadow-red-600"
            >
              [ Upload ]
            </button>
          </form>
        </div>
      }

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
