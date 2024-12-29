# Dockerfile
FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o user-service cmd/user-service/main.go

EXPOSE 8080
CMD ["./user-service"]
