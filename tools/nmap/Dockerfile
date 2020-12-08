FROM alpine:latest
RUN apk update \
    && apk add curl \ 
    nmap \
    nodejs npm\
    nmap-scripts \
    --no-cache bash \
    && mkdir nmap \
    && npm install xml2js
COPY . ./nmap
WORKDIR /nmap