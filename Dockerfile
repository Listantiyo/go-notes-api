FROM golang:1.26.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    GOOS=linux CGO_ENABLED=0 go build -o main ./cmd/server



FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

RUN mkdir /app/database

EXPOSE 8080

CMD ["./main"]