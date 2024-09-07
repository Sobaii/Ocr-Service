FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /ocr-service ./cmd/server

FROM alpine:latest

COPY --from=builder /ocr-service /ocr-service

EXPOSE 8080

CMD ["/ocr-service"]