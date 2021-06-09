#! /bin/bash
fname=$1
output=$2

result=$(grep -Ev "^#|Status: (Up|Down)" $fname \
        | cut -d' ' -f2,4- \
        | sed -n -e 's/Ignored.*//p' \
        | awk '{for(i=2; i<=NF; i++) { a=a""$i; }; split(a,s,","); for(e in s) { split(s[e],v,"/"); printf "\"%s://%s:%s\"\n" , v[5], $1,v[1]}; a="" ;}' \
        | sed -e 's/?//' -e 's/$/,/' \
        | grep -E "([^.]|^)([0-9]{1,2}|1[0-9]{2}|2[0-4][0-9]|25[0-5]).([0-9]{1,2}|1[0-9]{2}|2[0-4][0-9]|25[0-5]).([0-9]{1,2}|1[0-9]{2}|2[0-4][0-9]|25[0-5]).([0-9]{1,2}|1[0-9]{2}|2[0-4][0-9]|25[0-5])([^.]|$)" \
        | grep -v 'unknown\|alt\|sun' \
        | grep 'http\|ssl\|netbios\|ssh\|smtp\|mysql\|rdp\|mongodb\|ncp\|memcached\|irc\|postgres\|rtsp\|sip\|telnet\|vnc\|xmpp' \
        | sed -e 's/ssl|http/https/g' -e '$ s/.$//' -e '1s/"/[&/' -e '$s/$/]/' > $output)
#        | grep -E "(25[0-5]|2[0-4][0-9]|[1]?[1-9][1-9][1-9]?)\.(25[0-5]|2[0-4][0-9]|[1]?[1-9][1-9][1-9]?)\.(25[0-5]|2[0-4][0-9]|[1]?[1-9][1-9][1-9]?)\.(25[0-5]|2[0-4][0-9]|[1]?[1-9][1-9][1-9]?)" > $output)