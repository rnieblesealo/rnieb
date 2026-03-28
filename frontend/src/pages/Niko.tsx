const Niko = () => {
  return (
    <div className="w-sm flex flex-col items-center justify-center mt-8">

      <span className="mb-8">404</span>

      {/* inner iframe dimensions should match niko game initwindow */}
      <iframe
        src="/niko/niko.html"
        style={{
          width: "100%",
          aspectRatio: "1/1"
        }}
      />

    </div>
  )
}

export default Niko 
