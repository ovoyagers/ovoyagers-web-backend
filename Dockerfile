# Stage 1: Build the Go binary
FROM golang:1.23.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download dependencies only when the mod files change
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -o jamme-app .

# Stage 2: Create a minimal image with just the Go binary
FROM gcr.io/distroless/static

WORKDIR /app

# Copy the statically linked Go binary from the builder stage
COPY --from=builder /app/jamme-app .
COPY --from=builder /app/.env .

# Run the Go binary
CMD ["./jamme-app"]
