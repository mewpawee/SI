#!/bin/bash

##### Constants
IMAGE="$1"
TAG="$2"
REGISTRY="registry.digitalocean.com/sec"
##### Main
doctl registry login
docker tag $IMAGE:$TAG $REGISTRY/$IMAGE:$TAG
docker push $REGISTRY/$IMAGE:$TAG
