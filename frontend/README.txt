 _ __ _ __ (_) ___| |__  
| '__| '_ \| |/ _ \ '_ \ 
| |  | | | | |  __/ |_) |
|_|  |_| |_|_|\___|_.__/ 
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

=== QUESTIONS ============================================================================

[ ] cookies?
[ ] nginx
[ ] xss attack?
[x] delete keyword in typescript?
[x] localstorage?
[ ] at a low level how does routing work?
[ ] what is fast refresh ( react )?
[ ]

=== BRAINSTORM / SPEC ====================================================================

stuff i wanna share brainstorm:
  * drawings ( kinda done already )
  * my favorite music ( spotify api )
  * programming projects
      displayed as icons kinda like my website 

      if you click on an icon it takes you to an article ( use markdown renderer ? )
  * skate clips/progress
      john hill trick list typa display

      each trick is marked as complete or not complete
        each completed trick adds some xp to my skater level
        ( so 3 cols: trick name | complete | xp )

      clicking on a given trick takes you to a page that displays 
      clips of me doing that trick

      below the john hill trick list there are 3 of my coolest highlights
  * random objects i own
      a minecraft inventory type ui with random pngs of cool objects i own
      ( or ones that mean a lot to me )

      maybe do a threejs view of my minecraft guy??? thisd be sick!

      hovering over an object pulls up minecraft-style UI that shows its name and desc
      maybe also have a rarity to show how cool it is :)
  * my music
      link bandcamp for now; figure out how to do a proper page for it later

=== GENERAL NOTES ========================================================================

* npm ci <---- npm install but stricter for ci environments
* client-max-body-size ( config ) <----- set max post size ( default is 1mb )

* --classic on snap disables sandboxing ( similar to pacman --disable-sandbox )
* certbot needs server_name ( the domain name! ) nginx config to be set
    this is how it knows what to issue a certificate for!
    
* ALPINE LINUX is a very VERY ( 5mb!!! ) distro that doesn't even have coreutils
  if we want certbot, we'll need something like ubuntu for some package managin'!

* docker containers aren't booted with systemd 

* delete keyword removes a property from an object 
  e.g. from {"john":1, "mary":2}, delete user.john would remove that property 

=== LOCALSTORAGE =========================================================================

simple key value store 
  scoped to origin ( e.g. localhost:3000 can't read what localhost:5173 stored )

=== UPGRADING TO HTTPS ===================================================================

certbot can request and generate our cert easily:
  $ certbot certonly --standalone -d <DOMAIN NAME> <--- spins up verification server
                                                        gets us our cert and privkey
  this requires port 80 to be free since http connection is required
    ( turn off the container! )

  whatever domain name we pass must already point to our ip!

  the certs will live in the vps and will be volumemounted to our frontend container

after we have the certs, we need to mount them to the container & configure nginx
  ( see nginx.conf; basically need to give it cert paths + configure http redirect )

  for this, our frontend container must expose port 443 ( https )

this should be pretty much it!

=== REACT ================================================================================

* {children} is of type React.ReactNode

CONTEXTS:
  help avoid constantly passing down state

  context itself = container for provider; holds no state
  context provider = what is in the container; holds actual state

    useContext is used to pull values from current provider

    *** i don't understand why we need both, why not just store the state in the context?
        WHY DO WE NEED A PROVIDER?

        ...i guess it's like the blueprint-instance pattern 
            context defines a blueprint ( set of values that context type provides ) 
            the provider is an instance of that blueprint with mutable state 

=== WORKING WITH WASM ====================================================================

TO COMPILE A C++ RAYLIB APP TO WASM:
    ( from my notes; might need some extra reading up on and just shows gen process )

  1. install emscripten ( brew install emsripten, what have you )
  2. COMPILE RAYLIB FOR WASM:
      a. clone raylib source
      b. go to /src
      c. change platform to web and set output dir to /web ( or whatever is convenient )
      d. run $ make PLATFORM=PLATFORM_WEB -B 

      we will get a .a file that is the wasm version of raylib
  3. POINT OUR GAME'S CMAKE TO THIS
      a. set raylib_DIR to wherever raylib is ( .local/ is a good, user-scoped place )
      b. target_link_libraries the raylib wasm object ( ...the .a object from step 1 )
      c. target_include_dirs the raylib headers
      d. set appropriate compiler flags
          don't forget to bundle the game data! browser has no idea of your filesystem
  4. BUILD AND RUN:
      a. FROM YOUR GAME, make a build folder for the web version ( if u want ); cd into it
      b. run $ emcmake cmake .. <--- to configure
      c. run $ emmake make      <--- to build
      
      you will get some .html, wasm files; they are ready 

      run $ npx serve <---- where .html is and run it in browser
        ( browser blocks loading wasm unless a server is running )

  5. present the game using an iframe 
