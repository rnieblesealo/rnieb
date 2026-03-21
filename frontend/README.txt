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

--- SPECS --------------------------------------------------------------------------------

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

--- QUESTIONS ----------------------------------------------------------------------------

[ ] cookies?
[ ] nginx
[ ] xss attack?
[x] delete keyword in typescript?
[x] localstorage?

--- NOTES --------------------------------------------------------------------------------

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

--- LOCALSTORAGE -------------------------------------------------------------------------

simple key value store 
  scoped to origin ( e.g. localhost:3000 can't read what localhost:5173 stored )
