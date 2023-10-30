# Use the official Golang image
FROM golang:latest

# Copy your Golang server code into the container
COPY . /app

# Set the working directory
WORKDIR /app

# Build the server binary
RUN go build -o server

# Expose the server's port
EXPOSE 8080

# Run the server
CMD ["/app/server"]
