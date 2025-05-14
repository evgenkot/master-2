FROM docker.io/library/golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./main.go

FROM docker.io/library/alpine:latest
WORKDIR /
COPY --from=builder /server /server
EXPOSE 8080
CMD ["/server"]
