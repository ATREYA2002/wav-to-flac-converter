# Use the official Golang image as a base
FROM golang:1.23.2

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage caching
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o wav-to-flac-converter main.go

# Expose the port that the service runs on
EXPOSE 8080

# Command to run the application
CMD ["./wav-to-flac-converter"]
