FROM golang:1.21-alpine3.17

COPY . /app

WORKDIR /app

RUN go build -o main

CMD ["./main"]