# Use the official Golang Docker image
FROM golang:1.20

# Set the working directory
WORKDIR /go/src

# Copy the required dependencies
COPY . .

# Build the application
RUN go build -o app

# Specify a command to run the application
CMD ["go", "run", "server.go"]