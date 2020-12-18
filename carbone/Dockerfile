FROM alpine:latest
RUN mkdir realPdf
COPY . / realPdf/
RUN apk add --update nodejs npm && apk add libreoffice && npm install carbone
