# Use the official Go image as the base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o bitcoin-notify .

# Expose port 8080 for the application
EXPOSE 8080

# Set the command to run the application when the container starts
CMD ["./bitcoin-notify"]
