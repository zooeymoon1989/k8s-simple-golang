FROM golang:alpine AS builder

# Create and change to the app directory.
WORKDIR /app
# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY . .

RUN go mod download

# Build the binary.
RUN go build -v -o server


FROM alpine

COPY --from=builder /app/server /app/server

WORKDIR /app

COPY app.env /app/app.env

EXPOSE 80

ENTRYPOINT ["/app/server"]