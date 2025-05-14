FROM docker.io/library/golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache protoc protobuf-dev
COPY proto proto
COPY server server
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./server

FROM docker.io/library/alpine:latest
WORKDIR /
COPY --from=builder /server /server
EXPOSE 50051
CMD ["/server"]
