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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sqli ./cmd

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Add CA certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set the working directory inside the minimal image
WORKDIR /root/

RUN ls -la

# # Copy the built binary from the builder stage
COPY --from=builder /app/sqli .

COPY .env .
# Expose port 5000 to the host
EXPOSE 5000

RUN ls -la
# Define the command to run the application
CMD ["./sqli"]

