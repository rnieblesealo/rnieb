FROM lopsided/archlinux:devel

# install dependencies
RUN pacman -Syu --noconfirm --disable-sandbox 
RUN pacman -S --noconfirm --disable-sandbox \
  go \
  gcc \
  pkg-config \
  zlib \
  imagemagick \
  libheif

# working directory inside container
WORKDIR /app

# NOTE: don't need copy directives for now since we're using volume mount

# copy go files
COPY go.mod go.sum ./
RUN go mod download

# copy all source code
COPY . .

# build & run
RUN go build -o rnieb-backend .
CMD ["./rnieb-backend"]
