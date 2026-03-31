const Niko = () => {
  return (
    <div className="w-sm flex flex-col items-center justify-center mt-8">
      {/* inner iframe dimensions should match niko game initwindow */}
      <iframe
        src="/niko/Squirrel.html"
        style={{
          width: "100%",
          aspectRatio: "1/1",
          zIndex: 999,
        }}
      />

    </div>
  )
}

export default Niko 
