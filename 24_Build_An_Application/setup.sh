#!/bin/bash

container_name="pg-container"

# Function to check if the database exists
db_exists() {
    docker exec "$container_name" psql -U postgres -lqt | cut -d \| -f 1 | grep -qw "$1"
}

# Pull the PostgreSQL image
docker pull postgres

# Stop and remove existing container if it exists
if docker inspect -f '{{.State.Running}}' "$container_name" &>/dev/null; then
    echo "Stopping and removing existing container..."
    docker stop "$container_name" && docker rm "$container_name"
fi

# Run the PostgreSQL container
docker run --name "$container_name" -e POSTGRES_PASSWORD=passwd -p 5432:5432 -d postgres

# Wait for PostgreSQL to start
echo "Waiting for PostgreSQL to start..."
until docker exec "$container_name" pg_isready -q -h localhost -p 5432; do
    sleep 1
done

# Try creating the database until it succeeds
while ! db_exists gotestdb; do
    echo "Creating database..."
    docker exec -ti "$container_name" createdb -U postgres gotestdb
    sleep 1
done

echo "Database created successfully."
