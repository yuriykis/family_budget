# Build stage
FROM golang:1.17.2-alpine AS build

# Set working directory
WORKDIR /app

# Install make
RUN apk add --no-cache make

# Copy the application source code and Makefile
COPY . .

# Build the application using the Makefile
RUN make build

# Final stage
FROM alpine:3.14.2

# Set working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/bin/user_service .
COPY --from=build /app/configs ./configs

# Expose port 8001
EXPOSE 8001

# Start the application
CMD ["./user_service",  "--config-path=/app/configs/server.yml"]
