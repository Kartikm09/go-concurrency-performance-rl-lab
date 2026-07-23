# Docker verification

Verified locally on 2026-07-23 with Docker Engine 29.5.2 on an ARM64 Linux virtual
machine. The image was rebuilt after the final source and lint fixes.

```console
$ docker compose up -d --build --wait
Container webhook-api Healthy
$ curl --fail --silent http://127.0.0.1:8083/health
{"status":"ok"}
$ docker compose ps
webhook-api   Up (healthy)   0.0.0.0:8083->8083/tcp
$ docker compose down
```

This verifies the documented local container workflow and health endpoint. It is not
a production deployment or a multi-architecture compatibility claim.
