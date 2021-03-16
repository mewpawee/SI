# DNSMAP

### Code Usage
sourcecode from: https://github.com/resurrecting-open-source-projects/dnsmap

```
sudo docker run \
-it \
--env targetDomain=${TARGET NAME} \
--env accessKey=${DO ACCESS KEY} \
--env secretKey=${DO SECRET KEY} \
--env spaceName=${DO SPACE NAME} \
--env resultFile=${RESULT FILE NAME} \
${IMAGE NAME} 
```
## How to use this thing
Scan example.com using a wordlist (-w /usr/share/wordlists/dnsmap.txt):

dnsmap example.com -w /usr/share/wordlists/dnsmap.txt



