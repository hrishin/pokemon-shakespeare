FROM golang:alpine AS build-env

ENV GO111MODULE=on \ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o pokemon ./cmd/main.go

# final stage
FROM alpine

WORKDIR /app

COPY --from=build-env /build/pokemon /app/

EXPOSE 5000

ENTRYPOINT ./pokemon