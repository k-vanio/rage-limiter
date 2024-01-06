FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o build/main ./cmd/main.go

COPY . .

EXPOSE 8080

CMD ["./build/main"]