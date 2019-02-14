FROM golang:1.8.1-alpine

# Add all source code
COPY ./server ./server/

# Indicate the binary as our entrypoint
ENTRYPOINT go run ./server/server.go

# Expose your port
EXPOSE 10001
