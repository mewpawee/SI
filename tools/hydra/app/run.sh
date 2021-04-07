#!/bin/sh
PROTOCOL_ENDPOINT=$1;

#set the delimeter
IFS=':'

#split the protocol
read -ra split_protocol_endpoint <<< "$PROTOCOL_ENDPOINT"
protocol=${split_protocol_endpoint[0]}

#set the delimeter back to default
IFS=' '

#create the protcol username:password file
bash dpl4hydra.sh $protocol

#save a vairable for this protocol
FILE=./protocols/dpl4hydra_$protocol.lst

#run the hydra
./hydra -C $FILE -t 1 $PROTOCOL_ENDPOINT 2>&1| tee /tmp/log 
