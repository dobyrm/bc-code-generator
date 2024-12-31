# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/bc-code-generator-app

# Run
CMD ["/bin/sh", "-c", "./build/bc-code-generator-app && tail -f /dev/null"]
