FROM alpine:latest
RUN apk add enscript && apk add ghostscript && mkdir forReport
COPY *.txt /forReport
#RUN enscript /forReport/*.txt --output=- | ps2pdf -sOwnerPassword="1234" -sUserPassword="1234" - > file.pdf
