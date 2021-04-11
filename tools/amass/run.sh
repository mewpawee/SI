#! /bin/bash
domain = $1
toNmap="tmp/amass/&1.txt"
amass enum -active -ip -o $toNmap -d domain