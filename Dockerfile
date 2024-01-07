FROM golang:1.20 as builder

RUN apt-get update && \
    apt-get install -y apache2-utils && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN chmod +x ./test_token_successs.sh
RUN chmod +x ./test_token_error.sh
RUN chmod +x ./test_ip_successs.sh
RUN chmod +x ./test_ip_error.sh

RUN go mod tidy

RUN go build -o build/main ./cmd/main.go

EXPOSE 8080

CMD ["./build/main"]