curl https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/common.txt -o wordlist.txt > /dev/null 2>&1
./dirb https://csi.cmkl.ac.th wordlist.txt -o /tmp/log
