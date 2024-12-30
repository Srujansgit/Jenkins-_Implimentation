# Single-stage Dockerfile
FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server
EXPOSE 9090
CMD ["./server"]

