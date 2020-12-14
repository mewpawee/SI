# Test

## dnsmap +
- your-endpoint - use only domain name without http or https tag or subdirectory e.g. <your-endpoint>/subfolder 

### submission
`argo submit -n argo test-scan.yaml -p tool=dnsmap -p arg="./dnsmap <your-endpoint> -r /tmp/dnsmap.txt"`

ex:
`argo submit -n argo test-scan.yaml -p tool=dnsmap -p arg="./dnsmap testphp.vulnweb.com -r /tmp/dnsmap.txt"`

## dirb +
- your end point should be in the format of  "http://host/" or "https://host"

### submission
`argo submit -n argo test-scan.yaml -p tool=dirb -p arg="dirb <your-endpoint> -f"`

ex:
`argo submit -n argo test-scan.yaml -p tool=dirb -p arg="dirb http://testphp.vulnweb.com -f -o /tmp/dirb.txt`

## nmap +
- should not include  "http://" or "https://"
use this options
- nmap -v -A scanme.nmap.org
- nmap -v -sn 192.168.0.0/16 10.0.0.0/8
- nmap -v -iR 10000 -Pn -p 80

### submission
`argo submit -n argo test-scan.yaml -p tool=nmap -p arg="nmap -v -A <your-endpoint> -f"`  

ex:
`argo submit -n argo test-scan.yaml -p tool=nmap -p arg="nmap -v -A testphp.vulnweb.com -oN /tmp/nmap.txt"`

## searchsploit  //not yet
* d

` asdf`

##  sqlmap +
- path must included .php file
- example: --url http://testphp.vulnweb.com/listproducts.php?cat=1 --dump-all

`argo submit -n argo test-scan.yaml -p tool=sqlmap -p arg="--url <your-endpoint> --dump-all"`

## xsssniper  
- your end point must include "http" or "https"

### submission
`argo submit -n argo test-scan.yaml -p tool=xsssniper -p arg="python xsssniper.py --random-agent --crawl --forms -u "<your-endpoint>" > /tmp/log"`

ex:
`argo submit -n argo test-scan.yaml -p tool=xsssniper -p arg="python xsssniper.py --random-agent --crawl --forms -u "https://web.ctflearn.com" > /tmp/log"`

## hydra
* d

### submission
`argo submit -n argo test-scan.yaml -p tool=hydra -p arg="hydra -L /inputs/username -P /inputs/password <protocal>://<your-endpoints> -t 4 > /tmp/log 2>&1 || true"`

ex:
`argo submit -n argo test-scan.yaml -p tool=hydra -p arg="hydra -L /inputs/username -P /inputs/password ssh://188.166.188.94 -t 4 > /tmp/log 2>&1 || true"`