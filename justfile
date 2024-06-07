# Define the root directory
project_root := justfile_directory()

# Directories
songs_directory := project_root / "songs"
scripts_directory := project_root / "scripts"
deployments_directory := project_root / "deployments"

# Commands
deno_command := "deno run --allow-read --allow-write --allow-net --allow-env"

docker_compose_file := deployments_directory / "docker-compose.yaml"
docker_compose_command := "docker compose -f " + docker_compose_file

# Minio service configuration
minio_web_port := "9000"
minio_api_port := "9001"
minio_root_user := "jorge123"
minio_root_password := "jorge123"
minio_buckets := "music-uploads,music-processed"

# Command to download songs
download-songs:
    @echo "Starting song download script..."
    SONGS_DIR={{songs_directory}} \
    WEB_ENDPOINT="https://mp3teca.co" \
    SERVER_ENDPOINT="https://severmp3teca.xyz/-/mp3" \
    {{deno_command}} {{scripts_directory}}/download-music.ts
    @echo "Song download script completed."

# Command to bring up services
start-services:
    @echo "Bringing up services..."
    WEB_PORT={{minio_web_port}} \
    API_PORT={{minio_api_port}} \
    MINIO_ROOT_USER={{minio_root_user}} \
    MINIO_ROOT_PASSWORD={{minio_root_password}} \
    MINIO_DEFAULT_BUCKETS={{minio_buckets}} \
    {{docker_compose_command}} up -d
    @echo "Services are up and running."

# Command to rebuild all services
[confirm("Are you sure you want to rebuild all services without cache? [Y/N]")]
rebuild-services:
    @echo "Stopping services..."
    {{docker_compose_command}} down
    @echo "Rebuilding and bringing up services without cache..."
    {{docker_compose_command}} up --build --no-cache -d
    @echo "Services have been rebuilt and are running."

# Additional useful Docker recipes

# Command to stop a single or multiple services
[confirm("Are you sure you want to stop these services? [Y/N]")]
stop-services *services:
    @echo "Stopping all services..."
    {{docker_compose_command}} down {{services}}

# Command to stop all services
[confirm("Are you sure you want to stop all services? [Y/N]")]
stop-all-services:
    @echo "Stopping all services..."
    {{docker_compose_command}} down
    @echo "All services have been stopped."

# Command to view the logs of all services
view-logs:
    @echo "Displaying logs for all services..."
    {{docker_compose_command}} logs -f

# Command to remove all unused Docker objects
cleanup-docker:
    @echo "Removing all unused Docker objects..."
    docker system prune -f
    @echo "Cleanup completed."

# Command to check the status of services
check-status:
    @echo "Checking the status of services..."
    {{docker_compose_command}} ps
    @echo "Status check completed."
