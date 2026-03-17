# same as other test but with heic image instead

TEST_IMG="./me.heic"

PORT="8080"
ENDPOINT="/upload"

curl -X POST http://localhost:"$PORT""$ENDPOINT" \
  -F "image=@$TEST_IMG"
