# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev make

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Download dependencies first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate swagger docs
RUN make swagger

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/build/books-api-go /app/cmd/api/main.go

# Final stage
FROM alpine:3.19

WORKDIR /app

# Add non root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/build/books-api-go .

# Copy wait-for-it script
COPY deployments/docker/scripts/wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

# Use non root user
USER appuser

EXPOSE 8080

ENTRYPOINT ["./books-api-go"]