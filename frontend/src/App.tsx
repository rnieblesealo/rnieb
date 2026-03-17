import axios from "axios"
import { useEffect, useState } from "react"

const Collage = () => {
  const [imageFilenames, setImageFilenames] = useState([])

  // fetch images
  useEffect(() => {
    axios.get("http://localhost:8080/list-images")
      .then(res => {
        setImageFilenames(res.data)
        console.log(res.data)
      })
      .catch(err => {
        console.error("Failed to fetch images:", err)
      })
  }, [])

  return (
    /* render the images */
    <div>
      {imageFilenames.map(imageFilename => (
        <img key={imageFilename} src={`http://localhost:8080/uploads/${imageFilename}`} />
      ))}
    </div>
  )
}

export default function App() {
  return (
    <>
      {/* upload image form */}
      <form
        action="http://localhost:8080/upload"
        method="POST"
        encType="multipart/form-data"
      >
        <input type="file" name="image" /> {/* puts this imagae under the "image" field; look at curl */}
        <button type="submit">Upload</button>
      </form>

      {/* display images */}
      <div>
        <Collage />
      </div>
    </>
  )
}
