#! /bin/bash
desGrep="/tmp/nmp2hydra-raw"
desXml="/tmp/nmap2db.xml"
desLog="/tmp/log"
desHyd="/tmp/nmap2hydra"
ipList="/mnt/dnsmap/dnsmap"

nmap -v -A -p- -oG "$desGrep" -oX "$desXml" -oN "$desLog" -iL $ipList
node app.js "$desXml"
./n2h.sh "$desGrep" "$desHyd"