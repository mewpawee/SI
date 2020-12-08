#! /bin/bash
fname=$1
output=$2
result=$(egrep -v "^#|Status: (Up|Down)" $fname | cut -d' ' -f2,4- | sed -n -e 's/Ignored.*//p' | awk '{for(i=2; i<=NF; i++) { a=a""$i; }; split(a,s,","); for(e in s) { split(s[e],v,"/"); printf "%s//%s:%s\n" , v[5], $1,v[1]}; a="" ;}' | grep -v "http" > $output)
echo "$result"