FROM lopsided/archlinux:devel

# working directory inside container
WORKDIR /app

# sync mirrors 
# update system packs
# install my deps

# run as one command; sy step might be cached which causes -s to fail 
# ( no sync took place )
RUN  \
  pacman -Sy --noconfirm --disable-sandbox && \
  pacman -Su --noconfirm --disable-sandbox && \
  pacman -S --noconfirm --disable-sandbox \
  go \
  gcc \
  pkg-config \
  zlib \
  imagemagick \
  openjpeg2 \
  dav1d \
  ffmpeg \
  rav1e \
  libheif \
  libx264 \ 
  sqlite

# copy go files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# copy all source code
COPY . .

# build & run
RUN go build -o rnieb-backend -v
CMD ["./rnieb-backend"]
