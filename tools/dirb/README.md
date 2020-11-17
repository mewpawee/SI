# DIRB

Build a dirb microservice from bitnami/minideb:latest

## Build

Build image

`sudo docker build . -t dirb:v1`

Run image

`sudo docker run -it dirb:v1`

## Environment variable
- wordlist
- target
- output


## Example Command
``` 
docker run \
-it \
--mount type=bind,source="$(pwd)"/mylists,target=/mylists \
--mount type=bind,source="$(pwd)"/outputs,target=/outputs \
--env wordlist=/mylists/small.txt \
--env target=http://webscantest.com \
--env output=/outputs/output.txt \
dirb:v1 
```