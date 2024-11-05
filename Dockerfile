# Use a lightweight base image with Go
FROM golang:1.16-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules and dependencies files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal image for production
FROM alpine:latest

# Set the working directory in the final container
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 8080

# Run the Go app
CMD ["./main"]
