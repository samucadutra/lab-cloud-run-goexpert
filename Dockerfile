FROM golang:latest as build
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cloudrun ./cmd/main.go

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=build /app/cloudrun .
COPY .env .
ENTRYPOINT ["./cloudrun"]
