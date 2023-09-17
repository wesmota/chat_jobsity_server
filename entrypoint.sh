#!/bin/bash

# Wait for PostgreSQL to be ready
until pg_isready -h db; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 1
done

go install github.com/pressly/goose/v3/cmd/goose@latest

# Run the Go application
make run 
