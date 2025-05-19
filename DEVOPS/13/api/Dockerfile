FROM docker.io/library/golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gin-api .

FROM docker.io/library/alpine:latest
WORKDIR /app
COPY --from=build /app/gin-api .
EXPOSE 8080
CMD ["./gin-api"]

