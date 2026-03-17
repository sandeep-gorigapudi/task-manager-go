# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server


# Runtime stage
FROM alpine:latest

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/main .

USER appuser

EXPOSE 8080

CMD ["./main"]