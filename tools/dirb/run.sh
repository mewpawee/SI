#!/bin/sh
curl https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/common.txt -o wordlist.txt > /dev/null 2>&1
./dirb $1 wordlist.txt -f -o /tmp/log
