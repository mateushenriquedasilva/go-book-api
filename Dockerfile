# Stage 1: Build the Go application
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dir/main cmd/main.go

# Stage 2: Create the final, minimal runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dir .

EXPOSE 8080
CMD ["./main"]