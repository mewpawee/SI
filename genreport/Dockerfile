FROM alpine:latest
RUN apk update
RUN mkdir result
COPY result.txt /result
RUN cd /result && grep -Eo "([0-9]{1,3}\.){3}[0-9]{1,3}" result.txt > /result/filtered-result