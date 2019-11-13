FROM golang:1-alpine AS builder

COPY . /app
WORKDIR /app

RUN go build -o grpc_test cmd/grpc_test/main.go

FROM alpine

COPY --from=builder /app/grpc_test /grpc_test
COPY --from=builder /app/web /web
WORKDIR /
CMD ["/grpc_test"]