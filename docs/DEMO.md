# Payment Platform Demo

Run these commands in the Linux VM from the repository root.

```bash
cd /path/to/payment-platform
```

## Build

```bash
go build ./...
go test ./... -v
docker build -f services/payment-service/Dockerfile -t payment-service:local .
docker build -f services/frontend/Dockerfile -t levintovich/payment-platform-frontend:latest .
```

## Start

```bash
docker compose up --build
```

Kong is available at `http://localhost:8000`.

## Health Checks

```bash
curl -i http://localhost:8000/healthz
curl -i http://localhost:8000/readyz
```

Both requests should return `200 OK`.

## API Demo

Create a payment:

```bash
curl -i -X POST http://localhost:8000/payments \
  -H 'Content-Type: application/json' \
  -d '{"reference":"ORDER-1001","amount":"99.99","currency":"USD"}'
```

Expected response: `201 Created`.

From Windows, replace `localhost` with the Linux VM IP address, for example:

```text
http://<VM_IP>:8000/healthz
```
