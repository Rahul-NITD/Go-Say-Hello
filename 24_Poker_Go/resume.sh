container_name="pg-container"
docker start "$container_name"
echo "Waiting for PostgreSQL to start..."
until docker exec "$container_name" pg_isready -q -h localhost -p 5432; do
    sleep 1
done
echo "Database started successfully."
