# Use the official Golang image
FROM golang:latest

# Copy your Golang client code into the container
COPY . /app

# Set the working directory
WORKDIR /app

# Build the client binary
RUN go build -o client

# Run the client when the container starts
CMD ["/app/client"]
