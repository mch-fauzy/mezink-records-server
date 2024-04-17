#!/bin/bash

# Wait for the database container to become available
dockerize -wait tcp://db:3306 -timeout 1m

# Run database migration
sql-migrate up

# Start the API server
./main