FROM alpine:latest
RUN mkdir report
COPY . / report/
RUN apk add --update nodejs npm && apk add libreoffice && npm install carbone
