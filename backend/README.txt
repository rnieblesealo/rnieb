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

=== QUESTIONS ============================================================================

multipart forms?
what are fileheaders?
mebibytes?
go.sum?
layer caching?
pkgconfig?

=== LEARNING NOTES =======================================================================

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

--- IMAGEMAGICK + IMAGICK ( GO BINDINGS ) ------------------------------------------------

  * imagemagick version 7 only works with golang imagick v3 bindings
    ( ensure you pull the v3 bindings if get missing includes )

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
