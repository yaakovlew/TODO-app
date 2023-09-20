FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]