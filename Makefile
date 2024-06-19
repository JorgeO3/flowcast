# Define the root configs
project_root := $(CURDIR)
go_root_path := gitlab.com/JorgeO3/flowcast

# Directories
songs_directory := $(project_root)/songs
scripts_directory := $(project_root)/scripts
deployments_directory := $(project_root)/deployments
cmd_directory := $(project_root)/cmd
migrations_directory := $(project_root)/migrations

# Commands
deno_command := deno run --allow-read --allow-write --allow-net --allow-env
docker_compose_file := $(deployments_directory)/docker-compose.yaml

# Postgres service configuration
postgres_host := "localhost"
postgres_port := 5432
postgres_user := "jorge123"
postgres_password := "jorge123"
postgres_db_name := "auth_service"

define POSTGRES_ENVS
export POSTGRES_PORT=$(postgres_port) \
POSTGRES_USER=$(postgres_user) \
POSTGRES_PASSWORD=$(postgres_password) \
POSTGRES_DB=$(postgres_db_name);
endef

# Adminer service configuration
adminer_port := 8080

define ADMINER_ENVS
export ADMINER_PORT=$(adminer_port);
endef

# Minio service configuration
minio_web_port := 9000
minio_api_port := 9001
minio_root_user := jorge123
minio_root_password := jorge123
minio_buckets := music-uploads,music-processed

define MINIO_ENVS
export MINIO_WEB_PORT=$(minio_web_port) \
MINIO_API_PORT=$(minio_api_port) \
MINIO_ROOT_USER=$(minio_root_user) \
MINIO_ROOT_PASSWORD=$(minio_root_password) \
MINIO_DEFAULT_BUCKETS=$(minio_buckets);
endef

# Auth service configuration
auth_app_name := "auth-service"
auth_http_host := "0.0.0.0"
auth_http_port := "4100"
auth_database_url := "postgresql://$(postgres_user):$(postgres_password)@$(postgres_host):5432/$(postgres_db_name)?sslmode=disable"
auth_version := "v1.0.0"
auth_log_level := "debug"
migration_dir := $(migrations_directory)/auth

define AUTH_ENVS
export APP_NAME=$(auth_app_name) \
HTTP_HOST=$(auth_http_host) \
HTTP_PORT=$(auth_http_port) \
DB_NAME=$(postgres_db_name) \
MIGRATIONS_PATH=$(migration_dir) \
PG_URL=$(auth_database_url) \
VERSION=$(auth_version) \
LOG_LEVEL=$(auth_log_level);
endef

# Encapsulate environment variables and Docker Compose command
define DOCKER_COMPOSE_CMD
$(POSTGRES_ENVS) \
$(ADMINER_ENVS) \
$(MINIO_ENVS) \
$(AUTH_ENVS) \
docker-compose -f $(docker_compose_file)
endef

.PHONY: auth-service
auth-service:
	@echo "Starting auth service..."
	@$(AUTH_ENVS) go run $(cmd_directory)/auth
	@echo "Auth service is up and running."

# Command to download songs
.PHONY: download-songs
download-songs:
	@echo "Starting song download script..."
	@SONGS_DIR=$(songs_directory) \
	WEB_ENDPOINT=https://mp3teca.co \
	SERVER_ENDPOINT=https://severmp3teca.xyz/-/mp3 \
	$(deno_command) $(scripts_directory)/download-music.ts
	@echo "Song download script completed."

# Command to bring up services
.PHONY: start-services
start-services:
	@echo "Bringing up services..."
	@$(DOCKER_COMPOSE_CMD) up -d 
	@echo "Services are up and running."

# Command to rebuild all services
.PHONY: rebuild-services
rebuild-services:
	@read -p "Are you sure you want to rebuild all services without cache? [Y/N] " confirm && [ $${confirm} = Y ]
	@echo "Stopping services..."
	@$(DOCKER_COMPOSE_CMD) down
	@echo "Rebuilding services without cache..."
	@$(DOCKER_COMPOSE_CMD) build --pull --no-cache
	@echo "Raising up services"
	@$(MAKE) start-services
	@echo "Services have been rebuilt and are running."

# Command to stop a single or multiple services
.PHONY: stop-services
stop-services:
	@read -p "Are you sure you want to stop these services? [Y/N] " confirm && [ $${confirm} = Y ]
	@echo "Stopping services..."
	@$(DOCKER_COMPOSE_CMD) down $(services)

# Command to stop all services
.PHONY: stop-all-services
stop-all-services:
	@read -p "Are you sure you want to stop all services? [Y/N] " confirm && [ $${confirm} = Y ]
	@echo "Stopping all services..."
	@$(DOCKER_COMPOSE_CMD) down
	@echo "All services have been stopped."

# Command to view the logs of all services
.PHONY: view-logs
view-logs:
	@echo "Displaying logs for all services..."
	@$(DOCKER_COMPOSE_CMD) logs -f

# Command to remove all unused Docker objects
.PHONY: cleanup-docker
cleanup-docker:
	@echo "Removing all unused Docker objects..."
	docker system prune -f
	@echo "Cleanup completed."

# Command to check the status of services
.PHONY: check-status
check-status:
	@echo "Checking the status of services..."
	@$(DOCKER_COMPOSE_CMD) ps
	@echo "Status check completed."
