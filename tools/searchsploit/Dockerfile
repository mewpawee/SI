FROM ubuntu:latest

RUN apt-get update \
    && apt-get install -y git \
    && git clone https://github.com/offensive-security/exploitdb.git /opt/exploitdb \
    && ln -sf /opt/exploitdb/searchsploit /usr/local/bin/searchsploit \
    && mkdir ss
COPY . ./ss
WORKDIR /ss
