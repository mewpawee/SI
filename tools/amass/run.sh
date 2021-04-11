#! /bin/bash
domain = $1
toNmap="tmp/amass/domain.txt"
amass enum -active -ip -o $toNmap -d domain
