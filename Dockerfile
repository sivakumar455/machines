# Stage 1: Build the Go app
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Stage 2: Create a minimal image with only the executable
FROM alpine:latest

WORKDIR /root/

# Copy the executable from the builder stage
COPY --from=builder /app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]