#! /bin/bash
domain=$1
toNmap="tmp/domain.txt"
amass enum -active -ip -o $toNmap -d $domain 2>&1 | tee /tmp/log
