FROM lopsided/archlinux:devel

# our working directory inside container is:
WORKDIR /app

RUN  \
  pacman -Sy  --noconfirm --needed archlinux-keyring --disable-sandbox && \
  pacman -Su  --noconfirm --disable-sandbox && \
  pacman -S   --noconfirm --disable-sandbox \
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

# breakdown of this RUN:
#
# update arch master keys ( CRUCIAL TO PREVENT SIGNING ERRORS )& sync package db 
# upgrade system packages 
# install our dependencies

# install go modules first for better docker layer caching:
COPY go.mod go.sum ./
RUN go mod download

# copy source: 
COPY . .

# build & run
RUN go build -o rnieb-backend -v
CMD ["./rnieb-backend"]
