FROM golang:1.15 AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY ./cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

FROM alpine:latest
COPY --from=builder /app .
CMD ["./main"]
