# Use an official Go runtime as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code from the host to the container
COPY . .

# Build the Go application
RUN go build -o intern .

# Set the entry point command for the container
CMD ["./intern"]
