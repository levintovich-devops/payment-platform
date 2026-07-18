FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY services/payment-service/cmd/payment-service ./services/payment-service/cmd/payment-service
RUN go build -o /out/payment-service ./services/payment-service/cmd/payment-service

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/payment-service /app/payment-service
EXPOSE 8080
ENTRYPOINT ["/app/payment-service"]
