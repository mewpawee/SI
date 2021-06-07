# Cyber-Security-Index

CSI is a cloud-based automated security assessment system which could scan on server endpoints and create an assessment report on sceduled

**Table of Contents**
- [Documentation](#documentation)
## Documentation
## Object Storage directories structure
```
csi 
 |- [company-name]
 .   |- [scan-name]
 .   . 	   |- [endpoint]
 .   .     .       |- log
     .     .       |   |- dirb
           .       |   |   - [protocol:port].log
    	           |   |   .
 	               |   |   .
 	               |   |   .
 	          	   |   |
                   |   |- hydra
   		   		   |   |   - [protocol:endpoint].log
           	   	   |   |   .
           	   	   |   |   .
           	   	   |   |   .
           	   	   |   |
                   |   |- nmap
           	   	   |   |   - [endpoint].json
           	   	   |   |   - [endpoint].log
           	   	   |   |
           	   	   |   |- xsssniper
           	   	   |       - [protocol:port].log
          	   	   |	   .
           	   	   |       .
           	   	   |       . 	
          	   	   |
     	   	   	   |- tmp
               	       - report.json
```
