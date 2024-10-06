FROM golang:1.21.7 AS builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o app

FROM alpine
COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]