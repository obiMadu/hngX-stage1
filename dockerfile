# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY app .

# Expose port 8080 for the web application to listen on
EXPOSE 80

# Command to run the Go web application
CMD ["./app"]
