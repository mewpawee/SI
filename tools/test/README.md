# Test

## dnsmap
* your-endpoint - use only domain name without http or https tag or subdirectory e.g. <your-endpoint>/subfolder 

`argo submit -n argo test-scan.yaml -p tool=dnsmap -p arg="./dnsmap <your-endpoint>"`

