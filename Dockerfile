# Build step
FROM golang:1.22.3 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

# Image step
FROM scratch

WORKDIR /app

COPY --from=build /app/api /app

EXPOSE 8080

CMD ["./api"]