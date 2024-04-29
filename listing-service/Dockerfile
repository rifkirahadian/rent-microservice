# Use the official Golang image as the base image
FROM golang:1.22-alpine AS builder

# Install gcc
RUN apk add --no-cache gcc musl-dev

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Enable CGO and build the Go application
RUN CGO_ENABLED=1 go build -o /go-app

# Use a lightweight Alpine image for the final container
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the built executable from the builder stage to the final container
COPY --from=builder /go-app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./go-app"]
