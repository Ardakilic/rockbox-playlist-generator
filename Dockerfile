FROM golang:1.24-alpine AS build

# Install necessary build tools
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application with optimization flags
# CGO_ENABLED=0: Pure Go implementation, no C dependencies
# -ldflags="-s -w": Strip debug information to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o rockbox-playlist-generator

# Use distroless as minimal base image to package the application
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian12:nonroot AS run

# Copy the binary from build stage
COPY --from=build /app/rockbox-playlist-generator /app/rockbox-playlist-generator

# Copy .env.example as reference (can be used to create a .env file)
COPY .env.example /app/.env.example

# Set working directory
WORKDIR /app

# Use non-root user for better security
USER nonroot:nonroot

# Set the entrypoint with environment variable support
ENTRYPOINT ["/app/rockbox-playlist-generator"]

# Default command args (can be overridden by docker run command)
# Use -path with the container's music mount point
# The limit will be taken from env vars if available
CMD ["-path", "/music"] 