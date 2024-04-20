#!/bin/bash

# Load environment variables from .env file
source .env.docker

# Wait for the database container to become available
dockerize -wait tcp://$DOCKER_DB_MYSQL_HOST:$DOCKER_DB_MYSQL_PORT -timeout 1m

# Run database migration
OPTIONS="-config=dbconfig.yml -env development"
sql-migrate up $OPTIONS

# Start the API server
./main
