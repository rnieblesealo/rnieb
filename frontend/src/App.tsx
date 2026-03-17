const Collage = () => {
  return (
    <div>

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
