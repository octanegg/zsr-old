# Build Stage
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o core cmd/*.go

# Final Stage
FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/core .

EXPOSE 8080
CMD ["./core"]