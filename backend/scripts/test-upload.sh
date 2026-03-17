TEST_IMG="./chiikawa.png"

PORT="8080"
ENDPOINT="/upload"

curl -X POST http://localhost:"$PORT""$ENDPOINT" \
  -F "image=@$TEST_IMG"

# if we want multiple images, we just send all under same name:
# -F "image=@image-0.png"
# -F "image=@image-1.png"
# -F "image=@image-1.png"
# ...

# when we access them via MultipartForm, we call req.MultipartForm.File["image"]
#   this returns the list of handles to all files we uploaded under the image field name
