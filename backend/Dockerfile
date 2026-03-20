FROM lopsided/archlinux:devel

# working directory inside container
WORKDIR /app

# update mirrors 
RUN pacman -Sy --noconfirm --disable-sandbox

# update system packages
RUN pacman -Su --noconfirm --disable-sandbox

# install my dependencies
RUN pacman -S --noconfirm --disable-sandbox \
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
  sqlite

# copy go files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# copy all source code
COPY . .

# build & run
RUN go build -o rnieb-backend -v
CMD ["./rnieb-backend"]
