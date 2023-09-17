#!/bin/bash

# Wait for PostgreSQL to be ready
until pg_isready -h db; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 1
done

# Create the database
PGPASSWORD=password psql -h db -U root -c "CREATE DATABASE jobsity"

# Check if psql command was successful
if [ $? -ne 0 ]; then
  echo "Error: Failed to create database and schema"
  exit 1
fi

echo "Database 'jobsity' created successfully."

# Dump the schema
pg_dump postgres://root:password@db:5432/${DBNAME} --schema-only --no-owner --file db/schema.sql

# Check if pg_dump command was successful
if [ $? -ne 0 ]; then
  echo "Error: Failed to dump schema"
  exit 1
fi

# Print a message
echo "Database setup completed successfully."


# Run the Go application
./main
