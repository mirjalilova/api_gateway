# Builder stage
FROM golang:1.22.1 AS builder

WORKDIR /app

# Copy all the source files into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway .

# Final stage
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/gateway .

# Copy the configuration files from the builder stage
COPY --from=builder /app/api/model.conf ./api/
COPY --from=builder /app/api/policy.csv ./api/

# Copy the environment file
COPY .env .env 

# Ensure the binary is executable
RUN chmod +x gateway

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./gateway"]
