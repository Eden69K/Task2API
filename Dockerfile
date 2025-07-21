FROM golang:1.24.3-alpine AS builder

WORKDIR /app
COPY . .

RUN apk add --no-cache git
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api-gateway ./cmd/app/

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/api-gateway .
COPY config.yml ./config.yml
COPY .env .

EXPOSE 8080

CMD ["./api-gateway"]