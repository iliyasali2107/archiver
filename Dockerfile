# Build stage
FROM golang:1.20 AS builder

# Enable go garbage collection
ENV GODEBUG=gccheckmark=1

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app with a statically linked binary for a smaller image size
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o myapp

# Production stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/myapp /app/myapp

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]