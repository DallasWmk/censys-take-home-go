# Use the official Python image base to start
FROM golang:1.18

# Set working dir to the /app dir
WORKDIR /app

# Copy all Go module files and download dependancies
COPY go.mod go.sum ./
RUN go mod download

# Copy application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Run the Go application
CMD sleep 5 && ./main
