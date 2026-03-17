SRC="./drawings"
DEST="./drawings-converted"

# fixes literal when not all extensions are matched, TODO: understand better 
shopt -s nullglob

# make sure target dirs exist
if [ -d "$SRC" ]; then
  if [ ! -d "$DEST" ]; then
    echo "\n[ CREATING $DEST... ]\n"

    mkdir "$DEST"
  fi

  echo "\n[ CONVERTING & COMPRESSING... ]\n"

  # convert all heic and jpeg to png
  # since heic is huge, use generous compression

  for img in "$SRC"/*.{heic,HEIC,jpg,JPG}; do
    filename=${img##*/}  

    # ##*/ ---> match longest prefix ending in / from the start
    # "img" itself contains full pathname; this only gives us the filename itself
    # e.g. "./drawings/img.png" ----> "img.png"

    filename_no_ext=${filename%%.*} # %%.* ---> strip shortest suffix ( extension )
 
    echo "converting $filename_no_ext..."

    magick \
      "$img" \
      -format png \
      -quality 50 \
      "$DEST"/"$filename_no_ext".png

      # the % matches the prefix to .heic (i.e. the original filename)
      # the rename here would be "image.heic" ---> "image.png"
  done

  echo "\n[ CROPPING... ]\n"

  # take all images we just created and crop them to 1:1 aspect in place

  for img in "$DEST"/*.png; do
    filename=${img##*/}  
    filename_no_ext=${filename%%.*} # %%.* ---> strip shortest suffix ( extension )

    echo "cropping $filename_no_ext..."

    # TODO: what is difference between resize and extent?

    mogrify \
      -resize 500x500 \
      -extent 1:1 \
      -gravity Center \
      -quality 50 \
      "$img"
  done

  echo "\n[ DONE! ]\n"
else
  echo "$SRC does not exist!"
fi

# mogrify and convert do same things, 
# but mogrify does in place whereas convert makes a copy 
