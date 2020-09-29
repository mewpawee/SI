FROM bitnami/minideb:latest

RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/offensive-security/exploit-database.git /opt/exploit-database
RUN ln -sf /opt/exploit-database/searchsploit /usr/local/bin/searchsploit
RUN cp -n /opt/exploit-database/.searchsploit_rc ~/
