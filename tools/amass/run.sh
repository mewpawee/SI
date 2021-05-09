#! /bin/bash
domain=$1
toNmap="tmp/domain.json"
amass enum -active -ip -json $toNmap -d $domain 2>&1 | tee /tmp/log
