# Usage
> Usage: nmap [Scan Type(s)] [Options] {target specification}

# EXAMPLES:
## scan all including OS, version, script scanning, and traceroute
    nmap -v -A scanme.nmap.org
## scan TCP port 80 and 443
    nmap -sT -p 80,443 192.168.0.0/16
## scan in secretly or two-way handshake (normally three-handshake)
    nmap -sS -p 80,443 192.168.0.0/16
## scan os 
    nmap -O 192.168.0.0/16
> SEE THE MAN PAGE (https://nmap.org/book/man.html) FOR MORE OPTIONS AND EXAMPLES
