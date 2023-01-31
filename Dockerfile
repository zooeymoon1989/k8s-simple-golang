FROM golang:alpine AS builder

# Create and change to the app directory.
WORKDIR /app
# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
COPY main.go ./

RUN go mod download

# Build the binary.
RUN go build -v -o server


FROM alpine

COPY --from=builder /app/server /app/server

EXPOSE 3000

ENTRYPOINT ["/app/server"]