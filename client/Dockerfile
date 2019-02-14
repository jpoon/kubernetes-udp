FROM golang:1.8.1-alpine

# Add all source code
COPY ./client ./client/

# Indicate the binary as our entrypoint
ENTRYPOINT go run ./client/client.go
