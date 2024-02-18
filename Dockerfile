FROM golang:1-alpine3.18 AS build
LABEL authors="PethumJeewantha"

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/server/

FROM alpine:latest

WORKDIR /app

RUN adduser -D -g '' appuser \
    && chown -R appuser:appuser /app

COPY --from=build /app/main .

USER appuser

EXPOSE 3200

CMD ["./main"]
