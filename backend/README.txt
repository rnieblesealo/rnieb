 _ __ _ __ (_) ___| |__  
| '__| '_ \| |/ _ \ '_ \ 
| |  | | | | |  __/ |_) |
|_|  |_| |_|_|\___|_.__/ 
| |__   __ _  ___| | _____ _ __   __| |
| '_ \ / _` |/ __| |/ / _ \ '_ \ / _` |
| |_) | (_| | (__|   <  __/ | | | (_| |
|_.__/ \__,_|\___|_|\_\___|_| |_|\__,_|
       ___________
      |.---------.|
      ||         ||
      ||         ||
      ||         ||
      |'---------'|
       `)__ ____('
       [=== -- o ]--.
     __'---------'__ \
    [::::::::::: :::] )
     `""'"""""'""""`/T\
                    \_/

for now this only consists of an image processing service:
1. we upload a heic image to server
2. server receives image ( written to disk )
3. server imagemagicks it to convert to jpeg + compress ( heic is huge! ) 
4. image is secured + displayed on frontend 

we will also resize the image such that its width is always 500 ( for storage purposes )
  no need to crop it to square with magick 
  ( we can non-destructively do this in frontend w/styling )

  e.g start as 1280 x 600

      1280 / 600 = 2.1333 aspect

	    resize relative to width:

	    target width / actual width = conv factor
      500/1280 = 0.390625
      
      1280 * 0.390625 = 500 
      600  * 0.390625 = 234.375

      new dimensions = 500 x 234     
        ( take floor of decimals; handled in int conversion  )

      500 / 234 = 2.1333 which matches original aspect!

http returns ERROR on errors ( http.Error ) 
returns JSON STRUCT WITH MESSAGE on success
  ...maybe this response should be standard

  -----> decided to use standardized response always

we are opening/closing DB a lot of times...
  should we just repass one connection started once?

=== QUESTIONS ============================================================================

[x] multipart forms?
[x] what are fileheaders?
[x] mebibytes?
[x] go.sum?
[x] layer caching?
[x] pkgconfig?
[x] lanczos filtering? -- an expensive quality-centric image resizing algo
[x] cors?

=== LEARNING NOTES =======================================================================

--- MULTIPART FORMS ----------------------------------------------------------------------

* essentially just data grouped by names
  if two things share the same name, they are grouped under that name as a list
* FILE HEADERS are metadata about each file
  we use this to open the file itself
* interface{} is used to define a catchall type

--- GOLANG -------------------------------------------------------------------------------

* go.sum stores checksums of project deps
  on first go get, checksum is computed and stored here
  if we do go mod download ( e.g. on cloning repo ) that dl is verified against checksum

--- HOMEBREW -----------------------------------------------------------------------------

* brew install --cask <--- homebrew-cask is an extension for gui apps
* init() function is like main() but for go modules; runs on import 
  don't use this to defer cleanups!

--- DOCKER -------------------------------------------------------------------------------

* docker image = blueprint
         container = running instance of image
         dockerfile extends on blueprint image to add our stuff
* docker volume mount allows easy devcontainering:
  $ docker run -it -v $(pwd):/app rnieb-backend bash
                      -----------
                      ^host  ^container
* ^C'ing out of container wipes filesystem changes not in dockerfile
  ( container fs is ephemeral )
* container port needs to be exposed if we want to use it normally:
  $ docker run -p 8080:8080 -it -v $(pwd):/app rnieb-backend bash 
               ------------
                  ^host ^container
* docker build runs EVERYTHING except CMD directive
* docker run DOES run the CMD directive if no other command is specified
  e.g.
      docker run rnieb-backend      <---- will run the CMD in dockerfile
      docker run rnieb-backend bash <---- will override CMD and run the shell instead
* LAYER CACHING: each step in dockerfile creates a layer
                 layers are cached so that future builds pull from that cache

                 if a step requires redoing a layer ( say copying source ),
                 we must rerun that layer and everything below it on a new build

                 *** layers are filesystem snapshots that store a diff of the previous

--- PKGCONFIG ----------------------------------------------------------------------------

tells compilers where to find libraries' HEADER FILES and what COMPILER FLAGS they need 

installing libraries that support it creates .pc file into standard location 
( /usr/lib/pkgconfig ?) 

pkg-config reads these files and outputs compiler flags and locations of libraries

  *** this is how imagick ( go bindings of imagemagick ) find the real C imagemagick impl 

--- IMAGEMAGICK + IMAGICK ( GO BINDINGS ) ------------------------------------------------

* imagemagick version 7 only works with golang imagick v3 bindings
  ( ensure you pull the v3 bindings if get missing includes )

--- CORS ---------------------------------------------------------------------------------

request goes out
server receives it
server attaches access allow origin header to response
browser receives response and checks this header 
if the header allows this origin to make request, browser loads response into javascript

--- PACMAN & LANDLOCK --------------------------------------------------------------------

pacman (arch) recently implemented sandboxing for package installs:
  pacman writes stuff to disk while downloading

  if a malicious item is written in this process ( e.g. from a malicious mirror ),
  pacman's sandboxing mechanism uses LANDLOCK ( a kernel feature ) 
  to RESTRICT THE LOCATIONS PACMAN MAY WRITE TO

  this effectively reduces the "blast radius" of a malicious package install
   
docker containers share the host kernel w/restrictions;
the landlock syscall may fail

as such, we must DISABLE SANDBOXING to get pacman to work in docker

this is still SECURE;
the container is already a layer of sandboxing on its own!

--- MYSQL --------------------------------------------------------------------------------

mysql is now MARIADB ( damn... )

mysql_install_db            <---- sets up mysql; run once on install
  --user=mysql              <---- do everything mysql related under this user
  --basedir=/usr            <---- where mysql lives on disk
  --datadir=/var/lib/mysql  <---- where all mysql data lives on disk

mysqld_safe &               <---- start mysql server

mysql -u root               <---- connect to mysql server as root

( didn't continue using, but have notes for fun! )

--- MEBI VS MEGABYTES --------------------------------------------------------------------

hardware makers used standard meaning of mega ( 10^6 )
software, however, uses powers of 2 ( because binary! )

a MEGABYTE ( MB ) corresponds to the HARDWARE-centric meaning ( a power of 10 )
a MEBIBYTE ( MiB ) corresponds to the SOFTWARE-centric meaning ( a power of 2 )

1MB   = 10^6 = 1,000,000 bytes
1MiB  = 2^20 = 1,048,576 bytes

a quantity in MiB is ~5% less than MB

in software, always assume you are using MEBI when asked for byte sizes

so if you have a 32MB form size limit,
you must convert to MiB by doing << 20 ( left shift is equal to * 2^20 )

e.g.
  32MB to mebibytes: 
    32 << 20 
  = 32 * 2^20 
  = 32MiB

in short:
  JUST LEFT SHIFT BY 20!

CONVERSIONS:

1KiB: 2^10
1MiB: 2^20
1GiB: 2^30

  ( conversion factor is 2^10 )
