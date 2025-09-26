# -------- Stage 1: build --------
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o arritech-api ./cmd

# -------- Stage 2: runtime --------
FROM alpine:latest
RUN apk add --no-cache ffmpeg
# create user before copying, then chown during copy
RUN adduser -D appuser
WORKDIR /app
# copy the built binary (note the path!)
COPY --from=builder /app/arritech-api /app/arritech-api

# ensure appuser can read/execute
RUN chown appuser:appuser /app/arritech-api
USER appuser

EXPOSE 8080
ENTRYPOINT ["/app/arritech-api"]