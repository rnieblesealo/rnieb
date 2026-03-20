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

This is a fullstack app built for me to share my drawings, music, 
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

The backend is written in Go ( and runs in arch btw... )
The frontend uses TypeScripted React with Tailwind and Vite to build; it's served via nginx.
Both are dockerized using compose.
Everything is hosted in an Ubuntu Hetzner VPS.

Take a look at their respective dirs to read my notes and mini specs on each!

Thank you for visiting :)

OVERVIEW OF WHAT I'VE LEARNED:
  ( see individual frontend and backend readmes for more... )

 _                _                  _ 
| |__   __ _  ___| | _____ _ __   __| |
| '_ \ / _` |/ __| |/ / _ \ '_ \ / _` |
| |_) | (_| | (__|   <  __/ | | | (_| |
|_.__/ \__,_|\___|_|\_\___|_| |_|\__,_|
     ____________________________
    /                           /\
   /                          _/ /\
  /                          / \/
 /                           /\
/___________________________/ /
\___________________________\/
 \ \ \ \ \ \ \ \ \ \ \ \ \ \ \

  [ golang ]

    * getting comfortable with writing/reading go
    * using go's patterns correctly ( defers, error handling... )
        creating/destroying stuff appropriately ( db driver, imagemagick instance )
    * the layout of a go project ( packages, modules ); using external deps
    * doing http with net/http
        writing and settning up middleware + handlers
        multipart form parsing
        encoding json w/custom structs 
    * doing sqlite/database stuff with go-sqlite3 and database/sql
    * auth with jwts
    * pathing/os stuff with os and filepath packages

  [ docker ]
    ( my first time properly using docker... )
  
    * conceptual basics 
      ( image vs container, layer caching, etc... )
    * practical basics
        ( docker buildx build, docker ps, docker image, etc... )
    * using volume mounts ( i <3 volume mounts!!! )
    * using port maps
    * running fricken arch linux in docker ( the backend runs in arch btw... ) 
    * multistage builds and how they make things SMOL
    * docker compose ( i fricken LOVE docker compose!!! )
      intercontainer networking by service name ( this is really cool )

  [ sqlite ]
    ( also my first time doing proper db stuff...! )

    * the conceptual basics ( why sqlite? it's a single file, it requires no conn, etc. )
    * the practical basics 
      ( creating table, queriying, inserting, deleting, renaming stuff... )
    * using db browser for sqlite ( gui browser )
    * using sqlite3 cli

      also did some mysql ( mariadb? ) 
        and learned a bit about it before switching to squeelite:

        * set up mysql and connect to it

  [ image processing with imagemagick / using a c api in go ]

    * the patterns of using c bindings within go
    * using magickwands correctly ( one per image! )
    * various processing stuff
        ( format conversion, resizing calculations, dealing with heic... )
    * making all of this into a service
    * pkgconfig
  
  [ arch linux/linux in general ]
    
    * landlock kernel feature and how pacman uses it to keep packinstalls secure
        ( pacman is the sickest package manager and it's not even close )

  [ ci/cd; github actions ]

    * setting up a simple deploy job ( ssh and deploy! )
      environment vs. repo secrets

  [ hosting on a hetzner server with a domain ]
 
    * setting up a hetzner vps with docker
    * setting up a subdomain to point to it
    * dealing with ssh keys
        known_hosts and authorized_keys
    * using recovery mode ( i suffered greatly... )
    * the realization that ipv6 is still pretty unsupported and using ipv4
      will save you years of life
    * doing normal ssh/hosting stuff in ipv6

  __                 _                 _ 
 / _|_ __ ___  _ __ | |_ ___ _ __   __| |
| |_| '__/ _ \| '_ \| __/ _ \ '_ \ / _` |
|  _| | | (_) | | | | ||  __/ | | | (_| |
|_| |_|  \___/|_| |_|\__\___|_| |_|\__,_|
      _________
    / ======= \
   / __________\
  | ___________ |
  | | -       | |
  | |         | |
  | |_________| |________________________
  \=____________/                        )
  / """"""""""" \                       /
 / ::::::::::::: \                  =D-'
(_________________) 

  * dockerizing an nginx frontend
  * nginx & configuring a container with it
      serving static files and handling actual routes ( $uri... )
      setting up backend proxy ( backend is also dockerized! )
  * <form>s and sending data over via POSTs!
  * using localStorage
  * setting axios default headers
