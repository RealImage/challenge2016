FROM golang:1.21

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main

ENTRYPOINT ["./main"]
