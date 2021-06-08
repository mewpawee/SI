#! /bin/bash
domain=$1
out="tmp/"

sqlmap --url $1 --dump-all --out-put $out --batch 