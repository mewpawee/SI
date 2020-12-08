# Test

## dnsmap +
- your-endpoint - use only domain name without http or https tag or subdirectory e.g. <your-endpoint>/subfolder 

`argo submit -n argo test-scan.yaml -p tool=dnsmap -p arg="./dnsmap <your-endpoint>"`

## dirb +
- your end point should be in the format of  "http://host/" or "https://host"

`argo submit -n argo test-scan.yaml -p tool=dirb -p arg="dirb <your-endpoint> -f"`

## nmap +
- should not include  "http://" or "https://"
use this options
- nmap -v -A scanme.nmap.org
- nmap -v -sn 192.168.0.0/16 10.0.0.0/8
- nmap -v -iR 10000 -Pn -p 80

`argo submit -n argo test-scan.yaml -p tool=nmap -p arg="nmap -v -A <your-endpoint> -f"`  

## searchsploit -
* d

` asdf`

##  sqlmap +
- path must included .php file
- example: --url http://testphp.vulnweb.com/listproducts.php?cat=1 --dump-all

`argo submit -n argo test-scan.yaml -p tool=sqlmap -p arg="--url <your-endpoint> --dump-all"`

## xsssniper +
* d

`python xsssniper.py -u <your-endpoint >`

## hydra +
* d

`python xsssniper.py -u <your-endpoint >`