# Use the official Golang image
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy all the files from the current directory to /app directory inside the container
COPY . .

# Install psql (PostgreSQL client)
RUN apt-get update && apt-get install -y postgresql-client make

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Make the entrypoint script executable
RUN chmod +x /app/entrypoint.sh

# EXPOSE and CMD as before
EXPOSE 8080
ENTRYPOINT ["/app/entrypoint.sh"]
