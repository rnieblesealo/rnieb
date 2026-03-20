            __       _     
 _ __ __ _ / _| __ _( )___ 
| '__/ _` | |_ / _` |// __|
| | | (_| |  _| (_| | \__ \
|_|  \__,_|_|  \__,_| |___/
     _          __  __ 
 ___| |_ _   _ / _|/ _|
/ __| __| | | | |_| |_ 
\__ \ |_| |_| |  _|  _|
|___/\__|\__,_|_| |_|  

⠀⠀⠀⠀⢀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⣠⠾⠛⠶⣄⢀⣠⣤⠴⢦⡀⠀⠀⠀⠀
⠀⠀⠀⢠⡿⠉⠉⠉⠛⠶⠶⠖⠒⠒⣾⠋⠀⢀⣀⣙⣯⡁⠀⠀⠀⣿⠀⠀⠀⠀
⠀⠀⠀⢸⡇⠀⠀⠀⠀⠀⠀⠀⠀⢸⡏⠀⠀⢯⣼⠋⠉⠙⢶⠞⠛⠻⣆⠀⠀⠀
⠀⠀⠀⢸⣧⠆⠀⠀⠀⠀⠀⠀⠀⠀⠻⣦⣤⡤⢿⡀⠀⢀⣼⣷⠀⠀⣽⠀⠀⠀
⠀⠀⠀⣼⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠙⢏⡉⠁⣠⡾⣇⠀⠀⠀
⠀⠀⢰⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠋⠉⠀⢻⡀⠀⠀
⣀⣠⣼⣧⣤⠀⠀⠀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⡀⠀⠀⠐⠖⢻⡟⠓⠒
⠀⠀⠈⣷⣀⡀⠀⠘⠿⠇⠀⠀⠀⢀⣀⣀⠀⠀⠀⠀⠿⠟⠀⠀⠀⠲⣾⠦⢤⠀
⠀⠀⠋⠙⣧⣀⡀⠀⠀⠀⠀⠀⠀⠘⠦⠼⠃⠀⠀⠀⠀⠀⠀⠀⢤⣼⣏⠀⠀⠀
⠀⠀⢀⠴⠚⠻⢧⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⠞⠉⠉⠓⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠈⠉⠛⠛⠶⠶⠶⣶⣤⣴⡶⠶⠶⠟⠛⠉⠀⠀⠀⠀⠀⠀⠀

Welcome!

This is a self-hosted fullstack app built for me to share my drawings, music, 
programming, and just about anything I want 
( and have the will to implement infrastructure/UI for! )

It's basically like a MySpace or an Instagram, 
but with MY rules, MY code, and MY style ;)

Maybe it's not as polished, modern, or functional as Instagram (or even MySpace!), 
but every single line of code was put there by me, 
and every part of this was made artisanally and with intention and love. 

It's really more of an art project than a programming one :)

It's also meant to be a learning exercise, 
because when I do things by hand and force myself to struggle, I learn *a lot* more! 

The backend is written in Go.
The frontend uses TypeScripted React with Tailwind and Vite to build.
Both are dockerized using compose.

Take a look at their respective dirs to read my notes and mini specs on each!

Thank you for visiting :)

-- LIST OF TOPICS/SKILLS I HAVE LEARNED ABOUT/USED FROM THIS --

FRONTEND/REACT

- Environment variables and base URL configuration in Vite
- React file serving and API calls without a base URL
- Favicon configuration
- TypeScript interfaces and types
- useState and useEffect
- Axios requests
- Form handling and preventDefault
- File uploads with FormData
- CORS
- Tailwind CSS
- Component composition and props

GO

- Package imports and shadowing
- Error handling
- Custom types and structs
- Function signatures with multiple return values
- Defer and init()
- Pointers and nullable arguments
- Type assertions
- Constants and package-level vars
- HTTP server with net/http
- Multipart form parsing
- JSON encoding
- SQLite with go-sqlite3
- JWT authentication
- Middleware chaining
- CGO and build tags
- Resolving home directory with os.UserHomeDir()
- Creating directories with os.MkdirAll()
- File server setup with http.Handle, http.StripPrefix, http.Dir
- Reading config from environment variables

IMAGEMAGICK / IMAGICK

- MagickWand initialization
- Reading and writing images
- Format conversion (HEIC to PNG)
- Resizing with aspect ratio
- Nearest neighbor vs Lanczos filtering
- Extent/crop to square
- pkg-config and CGO flags

SQLITE

- Setup and schema design
- SQL queries (SELECT, INSERT, DELETE)
- go-sqlite3 driver
- DB Browser for SQLite
- Deleting all rows from a table
- Renaming a column with ALTER TABLE

AUTH

- JWT tokens
- Signing and verification
- Auth middleware
- Cookie vs localStorage

DOCKER

- Images vs containers
- Dockerfile instructions
- Layer caching
- Volume mounts
- Port mapping
- Platform differences (ARM vs x86)
- Multi-stage builds and why they matter
- Running and executing containers
- Arch Linux in Docker
- pacman sandboxing issues
- npm ci vs npm install
- Building inside vs outside containers
- Executing a shell inside a running container with docker compose exec

DOCKER COMPOSE

- Service definitions
- Port mapping (host:container)
- Volume bind mounts for persistent data
- Inter-container networking by service name
- depends_on
- Rebuilding with --build

NGINX

- Reverse proxying to a backend container
- Serving static files
- try_files and why each argument exists
- client_max_body_size for file uploads
- Proxying API routes
- HTTPS with certbot

CI/CD

- GitHub Actions for automated deployment
- SSH deployment to a server
- GitHub secrets

DEPLOYMENT / NETWORKING

- SSH keys and how they work
- known_hosts and authorized_keys
- IPv4 vs IPv6
- DNS and subdomains
- Hetzner VPS setup
- Firewall rules

.    _    +     .  ______   .          .
  (      /|\      .    |      \      .   +
      . |||||     _    | |   | | ||         .
 .      |||||    | |  _| | | | |_||    .
    /\  ||||| .  | | |   | |      |       .
 __||||_|||||____| |_|_____________\__________
 . |||| |||||  /\   _____      _____  .   .
   |||| ||||| ||||   .   .  .         ________
  . \|`-'|||| ||||    __________       .    .
     \__ |||| ||||      .        handcoding 4evr <3
  __    ||||`-'|||  .       .    __________
 .    . |||| ___/  ___________             .
    . _ ||||| . _               .   _________
 _   ___|||||__  _ \\--//    .          _
      _ `---'    .)=\oo|=(.   _   .   .    .
 _  ^      .  -    . \.|
