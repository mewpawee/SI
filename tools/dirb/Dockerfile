FROM alpine AS build-env
RUN apk add --update --no-cache openssl wget curl-dev libcurl musl-dev make gcc curl
COPY ./dirb222 ./dirb222 
WORKDIR ./dirb222
RUN ./configure CFLAGS="-O2 -g -fcommon"
RUN make 
COPY run.sh run.sh

