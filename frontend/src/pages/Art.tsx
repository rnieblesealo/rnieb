import { useEffect, useState } from "react"
import Collage from "../components/Collage"
import useAuthContext from "../context/useAuthContext"
import type { Drawing, GetDrawingsResponse } from "../types"
import axios from "axios"

// use custom submission function to avoid json response page behavior

const Art = () => {
  const { loggedIn } = useAuthContext()
  const [drawings, setDrawings] = useState<Drawing[]>([])

  async function handleUploadForm(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault()

    // submit image upload request 

    const formData = new FormData(e.currentTarget)
    axios.post(`${import.meta.env.VITE_BASE_URL}/upload`, formData) // submit upload
      .then(() => {
        // on upload complete, refetch drawings to update display

        axios.get(`${import.meta.env.VITE_BASE_URL}/get-drawings`)
          .then(res => {
            const getDrawingsResponse: GetDrawingsResponse = res.data
            setDrawings(getDrawingsResponse.data)
          })
      })
  }

  useEffect(() => {
    // fetch drawings
    axios.get(`${import.meta.env.VITE_BASE_URL}/get-drawings`)
      .then(res => {
        const getDrawingsResponse: GetDrawingsResponse = res.data
        setDrawings(getDrawingsResponse.data)
      })
      .catch(err => {
        console.error("Failed to fetch images:", err)
      })

  }, [])

  return (
    <>
      {/* display images */}
      <div>
        <Collage
          drawings={drawings}
          loggedIn={loggedIn}
          setDrawings={setDrawings}
        />
      </div >

      {/* upload form */}
      {
        loggedIn &&
        <div className="flex flex-col items-center justify-center text-red-500 mt-4">
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
    </>
  )
}

export default Art
