FROM golang:1.24-alpine AS build

ENV GOPATH=/

WORKDIR /src

COPY ./ ./

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o data-aggregation-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/configs

COPY --from=build /src/configs . 

WORKDIR /root/postgres/migrations

COPY --from=build /src/internal/repository/postgres/migrations . 

WORKDIR /root/app

COPY --from=build /src/data-aggregation-service . 

CMD ["./data-aggregation-service"]