FROM golang:1.21 as build

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./bin/auth ./cmd/auth

FROM scratch

WORKDIR /app

COPY --from=build /app/bin/auth ./bin/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY ./config/ ./config/

CMD ["./bin/auth"]