FROM alpine:latest
RUN apk update && apk add curl && apk add nmap && apk add nmap-scripts