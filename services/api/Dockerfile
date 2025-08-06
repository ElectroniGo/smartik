FROM golang:1.24-bookworm AS build

WORKDIR /src

COPY services/api/go.mod services/api/go.sum ./
RUN go mod download

COPY services/api/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api


FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
  ca-certificates \
  && rm -rf /tmp/* \
  && rm -rf /var/lib/apt/lists/* \
  && groupadd -r -g 1001 golang \
  && useradd -r -u 1001 -g golang api

COPY --from=build --chown=api:golang /api /usr/local/bin/api

USER api

EXPOSE 1323

ENTRYPOINT ["api"]