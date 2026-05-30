FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/app

FROM alpine:3.21

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/server .

RUN chown appuser:appgroup ./server

USER appuser

EXPOSE 8080

CMD ["./server"]
