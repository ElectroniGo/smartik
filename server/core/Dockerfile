FROM golang:1.24-bookworm AS build

RUN apt-get update && apt-get install -y \
  tesseract-ocr \
  libtesseract-dev \
  libleptonica-dev \
  tesseract-ocr-eng \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app .

FROM debian:bookworm-slim

ARG PORT

RUN apt-get update && apt-get install -y \
  tesseract-ocr \
  tesseract-ocr-eng \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/* \
  && apt-get clean

RUN useradd -r -u 1001 appuser


COPY --from=build --chown=appuser:appuser /app /usr/bin/app

USER appuser

ENV PORT=${PORT}

EXPOSE 8080

ENTRYPOINT [ "app" ]