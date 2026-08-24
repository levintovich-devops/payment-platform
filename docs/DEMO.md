# Payment Platform Demo

Run these commands in the Linux VM from the repository root.

```bash
git clone https://github.com/levintovich-devops/payment-platform.git
cd payment-platform
```

## Deploy and Start

The payment-service image is `payment-platform-payment-service:latest`.

```bash
docker compose up --no-build
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

Stop the deployment:

```bash
docker compose down
```

From Windows, replace `localhost` with the Linux VM IP address, for example:

```text
http://<VM_IP>:8000/healthz
```
