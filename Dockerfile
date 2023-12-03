# Build stage
FROM golang:1.21-alpine3.18 AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/api/main.go

# Final stage
FROM alpine:3.18.4
WORKDIR /app
COPY --from=builder /app/main /app
COPY --from=builder /app/.env /app

EXPOSE 8080 9090
CMD [ "/app/main" ]