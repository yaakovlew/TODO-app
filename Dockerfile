FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o task ./cmd/main.go

CMD ["./task"]