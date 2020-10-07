# Build Stage
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o zsr cmd/*.go

# Final Stage
FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/zsr .
COPY --from=builder /app/docs/index.html ./docs/index.html

EXPOSE 8080
CMD ["./zsr"]