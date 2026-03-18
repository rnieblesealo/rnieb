import axios from "axios"
import { useEffect, useState } from "react"

const Collage = () => {
  const [imageFilenames, setImageFilenames] = useState([])

  // fetch images
  useEffect(() => {
    axios.get("http://localhost:8080/list-images")
      .then(res => {
        setImageFilenames(res.data)
      })
      .catch(err => {
        console.error("Failed to fetch images:", err)
      })
  }, [])

  return (
    /* render the images */
    <div className="grid grid-cols-3 w-128">
      {imageFilenames && imageFilenames.map(imageFilename => (
        <img
          key={imageFilename}
          src={`http://localhost:8080/uploads/${imageFilename}`}
          className="w-xs aspect-1/1 object-cover"
        />
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
        />
        <input
          type="text"
          name="description"
          placeholder="Description..."
        />
        <input
          id="upload-images"
          type="file"
          name="file"
          accept="image/*"
          className="w-fit"
        />
        <button type="submit" className="bg-red-500 text-black p-4 w-fit">Upload</button>
      </form>
    </div>
  )
}
