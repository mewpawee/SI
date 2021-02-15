#!/bin/sh
PROTOCOL=$1;
ENDPOINT=$2;

#create the protcol username:password file
bash dpl4hydra.sh $PROTOCOL;

#save a vairable for this protocol
FILE=./protocols/dpl4hydra_$PROTOCOL.lst;

./hydra -C $FILE -t 1 localhost $PROTOCOL > result.txt;