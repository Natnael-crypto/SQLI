# Stage 1: Build the Go application
FROM golang:1.23.2 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the dependency files (go.mod and go.sum) to the working directory
COPY go.mod go.sum ./

# Download the module dependencies
RUN go mod download

# Copy the entire source code into the working directory
COPY . .

# Build the Go application with the binary named `sqli`
RUN go build -o sqli ./cmd

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Add CA certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set the working directory inside the minimal image
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/sqli .

# Expose port 5000 to the host
EXPOSE 5000

# Define the command to run the application
CMD ["./app/sqli"]
