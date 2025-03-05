# Start from a Go base image matching your version
FROM golang:1.23 as builder

# Set working directory
WORKDIR /app

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
