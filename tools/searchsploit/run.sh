#! /bin/bash
dx=$1
tmptxt=$(grep "version.*port" $1 \
    | cut -d'"' -f 14,16 \
    | sed 's/"/ /' \
    | awk '{for (i=1;i<=NF;i++) if (!a[$i]++) printf("%s%s",$i,FS)}{printf("\n")}')
while IFS= read -r line; do
    searchsploit -wj $line > tmp.json
done <<< $tmptxt