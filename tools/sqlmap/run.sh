#! /bin/sh
endpoint=$1
protocolEndpoint=$2
out="./result"

python sqlmap.py --url $protocolEndpoint --dump-all --output-dir $out --batch 
if [ -s ./result/$endpoint/log ]
then
    cp ./result/$endpoint/log /tmp/log
fi
