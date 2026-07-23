FROM golang:1.26.5-alpine AS build
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM alpine:3.23
RUN addgroup -S app && adduser -S -G app -u 10001 app
COPY --from=build /out/api /usr/local/bin/webhook-api
USER app
ENV APP_PORT=8083
EXPOSE 8083
CMD ["webhook-api"]
