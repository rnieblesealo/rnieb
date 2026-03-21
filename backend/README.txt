 _ __ _ __ (_) ___| |__  
| '__| '_ \| |/ _ \ '_ \ 
| |  | | | | |  __/ |_) |
|_|  |_| |_|_|\___|_.__/ 
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

since it'll just be me admindashing, i'm just gonna have a single "admins" table 
  *** how will i secure the database passwords?
  *** autoincrement id's probably not good here since probing attack might happen

*** SQLITE IS NOT THREAD SAFE!
    find a way to handle concurrent requests 
    ( are 2 handlers running at same time concurrent ? )

login currently uses localstorage to cache creds
  this is vulnerable to xss attacks
  upgrading to http only cookie is the move --- do this later to first learn basic auth

now we really need to figure out filepathing stuff...

we should really upgrade to AS BUILDER pattern 

nginx frontend:
  since /api/ is proxied to backend, backend routes must be prefixed with /api/

=== QUESTIONS ============================================================================

[ ] authorized_keys + known_hosts?
[ ] http.handle vs http.handler? ( go ) / protected routes
[ ] interfaces in go
[ ] mitm error ???
[ ] ssh keys + server?
[x] AS builder dockerfile pattern ( intermediate build container )
[x] OPTIONS header?
[x] cors?
[x] gai.conf?
[x] go.sum?
[x] if ip is intrinsic to machine how can we reassign them in hetzner dash? -- is vps :p
[x] indirect in go.mod?
[x] jwts? claims?
[x] lanczos filtering? -- an expensive quality-centric image resizing algo
[x] layer caching?
[x] mebibytes?
[x] multipart forms?
[x] pkgconfig?
[x] resolv.conf?
[x] serve mux?
[x] what are fileheaders?

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
* go build -v <---- -v flag shows build output
                     useful for builds that take a while on server 
* // indirect ( go.mod ) <---- means this dep isn't imported by us in our code, 
                               but by another 3rd party dependency

* servemux is the thing that maps routes to handles
    we can either make our own by hand ( NewServeMux ) 
    and add functions to it specifically

    or use DefaultServeMux ( the one go http creates for us )

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

* docker image prune <---- removes <none> dangling images
* DOCKER COMPOSE ( docker-compose.yml ) helps define how to run containers

* docker containers may reach each other via servicename, no port expose required
  so a frontend container can talk to the backend

* AS builder makes an intermediate container that builds the thing
  we then copy to the real container in last stage  

  this is useful bc if we build in the final container, leftovers from build 
  ( source code, node modules, etc.) make the filesize huge when we only need the binary

* docker compose up --build       <---- rebuild this time!
* docker compose up <servicename> <----- only compose one of the services

--- PKGCONFIG ----------------------------------------------------------------------------

tells compilers where to find libraries' HEADER FILES and what COMPILER FLAGS they need 

installing libraries that support it creates .pc file into standard location 
( /usr/lib/pkgconfig ?) 

pkg-config reads these files and outputs compiler flags and locations of libraries

  *** this is how imagick ( go bindings of imagemagick ) find the real C imagemagick impl 

--- IMAGEMAGICK + IMAGICK ( GO BINDINGS ) ------------------------------------------------

* imagemagick version 7 only works with golang imagick v3 bindings
  ( ensure you pull the v3 bindings if get missing includes )

--- CORS / OPTIONS -----------------------------------------------------------------------

before all nonsimple requests, an OPTIONS preflight request is sent
  ( e.g. if i do POST /api/ping, OPTIONS /api/ping is sent first )
  
the preflight request contains info about the "real" request we're trying to make
e.g. "Origin" ( our url ) and "Access-Control-Request-Method" ( post? get? )

via cors middleware, the server sends back a response with what it allows
e.g. the notorious "Access-Control-Allow-Origin" ( the places that it allows to reach it )
  *** THIS RESPONSE MAY VARY BASED ON OUR OPTIONS INFO!
      the server may change what it allows based on the request origin or method, 
      for example
      
the client gets this response and checks it
if the server forbade anything that our request needs to go out 
the real request isn't allowed to leave

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

--- AUTHENTICATION -----------------------------------------------------------------------

* generate jwt secrets using ---> openssl -rand -base64 32 
    ( 32 bytes is std. length for HS256 secret encoding )

HOW A JWT WORKS:
  we define a secret ( 32 bytes is recommended but can be anything )
  
  we submit some creds to server ( signup )

  server issues a jwt with claims ( claims = data accompanying the jwt ) such as:
    * when the token expires
    * the username that the jwt is for 

  server sends the jwt SIGNED ( but not encrypted! ) using the secret back to client
    ( base64 decoding a jwt allows us to see the claims -- CAREFUL WITH WHAT THEY ARE! )

  client saves it somewhere

  client sends the jwt under the "Authorization" header 
  anytime they'd like to access something auth-protected

  server receives the request and tries to decode the jwt with auth middleware
  it also checks the claims to ensure it's valid ( e.g. is the expiration past? )

  if this fails, the middleware aborts and the main handler isn't run 
  ( we aren't authorized; issue a 401 )

  if this succeeds, the main handler does run

  NOTES:
    * RFC 7519 defines standard claims which jwt parsers know to look for
      "exp" ( expiration date ) and "sub" ( subject; the user id ) are notable
      in our use case
 
--- HETZNER HOSTING ----------------------------------------------------------------------

honestly i just picked them cuz they don't say ai anywhere in their homepage

to get ip:
  dashboard > servers > public ip ( it shouldn't trail with a /xx ...)

setup:
  ssh in ( make sure to set up keys )
  apt-get update
  apt-get install -y docker.io        <----------- docker setup
  systemctl start docker
  systemctl enable docker
    ( or just what u know: systemctl enable --now docker )
  
  clone ur code

  build the container

  run the container

NOTES:
  i tried ipv6 only setup

  this failed because most of the internet still uses ipv4
    AN IPV6 SERVER MAY ONLY REACH IPV6 THINGS

  i got mitm error when rebuilding server after switching to ipv4
    remove known host signature and try again

  ssh -v to view log

  root's home dir is at / level ( /root/ )

  because go build must freshly compile imagemagick, this will take a WHILE on server

TIPS:
  use scp -6 for ipv6 addresses to be parsed correctly
    use quotes around root@... part and wrap address in brackets,
      e.g. scp -6 -i rnieb/.ssh/id_ed25519 -r ./rnieb "root@[xxxx:xxxx:xxxx...]:/root/"

  flip scp order to copy from remote to local ( duh... )

--- DOMAIN STUFF ( DNS ) -----------------------------------------------------------------

RECORD TYPES:
  A           maps domain -> ip 
  CNAME       maps domain -> domain

HOST MEANINGS:
  @                   root domain ( e.g. rnieb.dev )
  www, stuff, api...  prefixes to root domain ( www.rnieb.dev, stuff.rnieb.dev... )

remember some ports are reserved for stuff ( thank u beej! )
  80          http
  443         https

dig <domain name> <----- checks if dns has propagated yet, USEFUL AF!

TRR = trusted recursive resolver
  if a domain has this true, it's resolved with a trusted entity ( cloudflare ) 
  instead of isp directly

  isp can't see dns resolution info; this helps do things more privately

gai.conf <----- defines how to prioritize ips when a hostname resolves to multiple ips
                ( e.g. do i prefer the v4 or v6 addr? )

resolv.conf <-- tells the system where to send dns queries
                this is typically not edited by hand and 
                managed by the system 

the server you are running is a VPS ( virtual private server )
  this means it's a vm, and so things that would be intrinsic to a physical machine
  ( e.g. an ip address ) can be changed/reassigned 
  ( they aren't defined by physical hardware )

--- CI/CD --------------------------------------------------------------------------------

github actions go under .github/workflows/

REPO SECRET <---- available to all workflows
ENVIRONMENT SECRET <----- scoped to specific environment ( we can give jobs scopes )

add workflow_dispatch to add manual action trigger button

need to add passphrase as well
  don't use these for cicd since the key itself is already secret; passphrase not needed
