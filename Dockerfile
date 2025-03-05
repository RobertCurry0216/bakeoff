# Start from a Go base image matching your version
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

RUN apt-get update && apt-get install -y curl


# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o bakeoff ./cmd/bakeoff

# Use a smaller base image for the final image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
ENV HUB_API_URL="http://app:3000"
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/bakeoff .
# Copy templates directory
COPY --from=builder /app/templates ./templates

# Expose the port the app runs on
EXPOSE 8081

# Command to run the executable
CMD ["./bakeoff"]
