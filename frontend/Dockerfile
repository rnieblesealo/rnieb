# build container
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --legacy-peer-deps
COPY . .
RUN npm run build

# final container
FROM ubuntu/nginx 
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

# certbot
# adapted from
# https://certbot.eff.org/instructions?ws=nginx&os=snap

# install snap ( not present in ubuntu/nginx )
RUN sudo apt update && sudo apt install snapd

# install certbot
RUN snap install --classic certbot

# make sure binary can be used systemwide
RUN ln -s /snap/bin/certbot /usr/local/bin/certbot

# have certbot edit nginx.conf to add https config
# ( see README to see what this'd look like )
RUN certbot --nginx
