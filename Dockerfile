FROM golang:1.22.6-alpine3.20 AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY resources resources

RUN CGO_ENABLED=0 GOOS=linux go build -o /garage ./cmd

FROM scratch

COPY --from=builder /garage /garage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/resources /resources

EXPOSE 8080

CMD ["/garage"]
