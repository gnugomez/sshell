# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder
# Install necessary packages (bash is needed to run menu.sh)
RUN apk add --no-cache bash git
WORKDIR /app
# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy all source files (main.go, menu.sh, host_key, etc.)
COPY . .
# Build a statically-linked binary
RUN CGO_ENABLED=0 go build -o sshell main.go

# Stage 2: Create a minimal runtime container
FROM alpine:3.18
# Install bash (required for your shell script)
RUN apk add --no-cache bash
# Create a non-root user and group for enhanced security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
# Copy the built binary, the menu script, and the persistent host key from the builder stage
COPY --from=builder /app/sshell .
# Switch to non-root user
USER appuser
# Expose the SSH server port
EXPOSE 2222
# Run the SSH server with required arguments
ENTRYPOINT ["/app/sshell"]

