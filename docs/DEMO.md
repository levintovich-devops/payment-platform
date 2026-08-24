# Payment Platform Demo

Run these commands in the Linux VM from the repository root.

```bash
cd /path/to/payment-platform
```

## Build

```bash
go test ./...
go build ./...
```

## Start

```bash
go run ./cmd/payment-service
```

The service listens on `http://localhost:8080`.

## Health Checks

```bash
curl -i http://localhost:8080/healthz
curl -i http://localhost:8080/readyz
```

Both requests should return `200 OK`.

## API Demo

Create a payment:

```bash
curl -i -X POST http://localhost:8080/payments \
  -H 'Content-Type: application/json' \
  -d '{"reference":"ORDER-1001","amount":"99.99","currency":"USD"}'
```

Copy the returned `id`, then capture the payment:

```bash
curl -i -X POST \
  http://localhost:8080/payments/<PAYMENT_ID>/capture
```

Expected responses: `201 Created`, then `200 OK`.

From Windows, replace `localhost` with the Linux VM IP address, for example:

```text
http://<VM_IP>:8080/healthz
```
