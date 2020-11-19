FROM golang:alpine
RUN mkdir /app
WORKDIR /app

COPY . .
RUN go build

EXPOSE 8080

CMD ["./backendAPI"]
