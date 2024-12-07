# Stage 1: Build the application
FROM golang:1.21 AS builder


# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Stage 2: Create the minimal image
FROM debian:bookworm-slim

# Copy the built binary from the builder stage
COPY --from=builder /app/main /main

# Copy the .env file into the container's root directory
COPY .env ./

# Ensure the binary is executable
RUN chmod +x /main

# Specify the entrypoint for the container
CMD ["/main"]