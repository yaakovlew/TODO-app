FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o task ./cmd/rest_app/main.go

CMD ["./task"]