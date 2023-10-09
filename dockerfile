# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy code to /app
COPY . /app

# Build the Go application
ENV GO111MODULE=auto
RUN go build -o app

# Expose port 8080 for the web application to listen on
EXPOSE 80

# Command to run the Go web application
CMD ["./app"]
