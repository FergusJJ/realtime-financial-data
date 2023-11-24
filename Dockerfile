FROM golang:1.20-bookworm as builder

WORKDIR /app

ADD . /app

RUN go mod tidy
RUN go build -o bin/server cmd/streamingservice/main.go

FROM debian:bookworm-slim

COPY --from=builder /app/bin/server .

EXPOSE 8000

CMD ["./server"]