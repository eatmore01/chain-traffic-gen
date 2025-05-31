FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/lalachka ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/lalachka .

CMD ["./lalachka"]