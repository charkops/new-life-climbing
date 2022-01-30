# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . .

# Release mode
# ENV GIN_MODE=release

RUN go mod download

CMD ["sh"]
