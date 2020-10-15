# Description
This service filter or clean the data that come from DNSMap service(s) and choose only nessessary data that we can use into another service. RIght now we choose *only* IPs that we going to use.
* This container dependency is result.txt that contain IPv4 addresses
* This is an example file **result.txt**
> classroom.kmitl.ac.th\
IP address #1: 161.246.35.53

> connect.kmitl.ac.th\
IP address #1: 10.252.23.104\
[+] warning: internal IP address disclosed\
IP address #2: 10.252.23.102\
[+] warning: internal IP address disclosed

> demo.kmitl.ac.th\
IP address #1: 161.246.35.92\
IP address #2: 161.246.35.92

> finance.kmitl.ac.th\
IP address #1: 161.246.34.137\
IP address #2: 161.246.34.137

> ftp.kmitl.ac.th\
IP address #1: 161.246.34.250

> hotspot.kmitl.ac.th\
IP address #1: 161.246.34.83

> hr.kmitl.ac.th\
IP address #1: 161.246.127.121

> ic.kmitl.ac.th\
IP address #1: 161.246.35.104

> imap.kmitl.ac.th\
IP address #1: 161.246.34.167\
IP address #2: 161.246.34.167

> la.kmitl.ac.th\
IP address #1: 161.246.35.133

> labs.kmitl.ac.th\
IP address #1: 161.246.35.223

> mail.kmitl.ac.th\
IP address #1: 161.246.34.12

> md.kmitl.ac.th\
IP address #1: 161.246.35.185

> ns1.kmitl.ac.th\
IP address #1: 161.246.52.21

> ns2.kmitl.ac.th\
IP address #1: 10.252.29.52\
[+] warning: internal IP address disclosed

> ntp.kmitl.ac.th\
IP address #1: 161.246.52.56

> pop.kmitl.ac.th\
IP address #1: 161.246.34.167\
IP address #2: 161.246.34.167

> qa.kmitl.ac.th\
IP address #1: 161.246.35.209

> se.kmitl.ac.th\
IP address #1: 161.246.35.104\
IP address #2: 161.246.35.104

> smtp.kmitl.ac.th\
IP address #1: 161.246.34.13

> staff.kmitl.ac.th\
IP address #1: 203.151.165.203

> unix.kmitl.ac.th\
IP address #1: 161.246.34.187

> web.kmitl.ac.th\
IP address #1: 161.246.35.98

> webmail.kmitl.ac.th\
IP address #1: 161.246.34.53

> www.kmitl.ac.th \
IP address #1: 161.246.34.11

> www2.kmitl.ac.th\
IP address #1: 161.246.35.98\
IP address #2: 161.246.35.98

> www3.kmitl.ac.th\
IP address #1: 161.246.34.35

* We select only available IPs in any domains and subdomains and not use any domain names that **do not** have IP(s)
* After run this container the result will be shown in directory in a named file as **filtered-result**
> /result/filtered-result
* This is an example result of **filtered-result** file
> 161.246.35.53\
10.252.23.104\
10.252.23.102\
161.246.35.92\
161.246.35.92\
161.246.34.137\
161.246.34.137\
161.246.34.250\
161.246.34.83\
161.246.127.121\
161.246.35.104\
161.246.34.167\
161.246.34.167\
161.246.35.133\
161.246.35.223\
161.246.34.12\
161.246.35.185\
161.246.52.21\
10.252.29.52\
161.246.52.56\
161.246.34.167\
161.246.34.167\
161.246.35.209\
161.246.35.104\
161.246.35.104\
161.246.34.13\
203.151.165.203\
161.246.34.187\
161.246.35.98\
161.246.34.53\
161.246.34.11\
161.246.35.98\
161.246.35.98\
161.246.34.35
