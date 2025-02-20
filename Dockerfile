# Development stage
FROM golang:1.23-alpine

WORKDIR /app

# Install build dependencies and Air
RUN apk add --no-cache \
    gcc \
    musl-dev \
    git \
    curl

# Install Air for hot reloading
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy Air configuration
COPY .air.toml ./

# Copy source code
COPY . .

EXPOSE 3000
# EXPOSE 2112

# Use Air for hot reloading
CMD ["air", "-c", ".air.toml"]
