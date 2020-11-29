FROM golang:alpine AS build-env

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o pokemon ./cmd/main.go


# final stage
FROM alpine

WORKDIR /app

COPY --from=build-env /build/pokemon /app/

EXPOSE 5000

ENTRYPOINT ./pokemon