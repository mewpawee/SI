#! /bin/bash
domain=$1
jsonForReport="tmp/domain.JSON"
toNmap="tmp/domain.txt"
amass enum -active -ip -o $toNmap -json $jsonForReport -d $domain 2>&1 | tee /tmp/log
