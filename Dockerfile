#docker buildx build --platform linux/amd64 .

FROM golang:1.21.1-alpine3.17 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

# EXPOSE 8000
# CMD ["/app/main"]

FROM alpine:3.13
WORKDIR /app

COPY --from=builder /app/main .
COPY ./email ./email
# COPY --from=builder /app/.env .

EXPOSE 8000
CMD ["/app/main"]
