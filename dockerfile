# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Install Git to clone the GitHub repository
RUN apk update && apk add --no-cache git

# Clone your Go application repository
RUN git clone https://github.com/obiMadu/hngX-stage1.git .

# Build the Go application (assuming it's in the current directory)
RUN export GO111MODULE=auto
RUN go build -o app

# Expose port 8080 for the web application to listen on
EXPOSE 80

# Command to run the Go web application
CMD ["./app"]