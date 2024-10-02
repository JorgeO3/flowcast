# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Define the root configs
project_root := $(CURDIR)
go_root_path := gitlab.com/JorgeO3/flowcast

# Directories
songs_directory := $(project_root)/songs
scripts_directory := $(project_root)/scripts
deployments_directory := $(project_root)/deployments
cmd_directory := $(project_root)/cmd
migrations_directory := $(project_root)/migrations
assets_directory := $(project_root)/assets

# Files
docker_compose_file := $(deployments_directory)/docker-compose.yaml
email_template_file := $(assets_directory)/email_template.html

# Commands
python_command := python
deno_command := deno run -A

# Catalog service
define CATALOG_SERVICE_ENVS
export CATALOG_APP_NAME=$(CATALOG_APP_NAME) \
CATALOG_HTTP_HOST=$(CATALOG_HTTP_HOST) \
CATALOG_HTTP_PORT=$(CATALOG_HTTP_PORT) \
CATALOG_DB_URL=$(CATALOG_DB_URL) \
CATALOG_DB_NAME=$(CATALOG_DB_NAME) \
CATALOG_VERSION=$(CATALOG_VERSION) \
CATALOG_LOG_LEVEL=$(CATALOG_LOG_LEVEL);
endef

# Catalog storage service
define CATALOG_STORAGE_ENVS
export CATALOG_STORAGE_PORT=$(CATALOG_STORAGE_PORT) \
CATALOG_STORAGE_ROOT_USER=$(CATALOG_STORAGE_ROOT_USER) \
CATALOG_STORAGE_ROOT_PASSWORD=$(CATALOG_STORAGE_ROOT_PASSWORD);
endef

define CATALOG_STORAGE_ADMIN_ENVS
export CATALOG_STORAGE_ADMIN_PORT=$(CATALOG_STORAGE_ADMIN_PORT) \
CATALOG_STORAGE_ADMIN_USER=$(CATALOG_STORAGE_ADMIN_USER) \
CATALOG_STORAGE_ADMIN_PASSWORD=$(CATALOG_STORAGE_ADMIN_PASSWORD) \
CATALOG_STORAGE_ADMIN_URL=$(CATALOG_STORAGE_ADMIN_URL) \
CATALOG_STORAGE_ADMIN_BASICAUTH=$(CATALOG_STORAGE_ADMIN_BASICAUTH);
endef

# Raw audio storage service
define RAW_AUDIO_STORAGE_ENVS
export RAW_AUDIO_STORAGE_WEB_PORT=$(RAW_AUDIO_STORAGE_WEB_PORT) \
RAW_AUDIO_STORAGE_API_PORT=$(RAW_AUDIO_STORAGE_API_PORT) \
RAW_AUDIO_STORAGE_ROOT_USER=$(RAW_AUDIO_STORAGE_ROOT_USER) \
RAW_AUDIO_STORAGE_ROOT_PASSWORD=$(RAW_AUDIO_STORAGE_ROOT_PASSWORD) \
RAW_AUDIO_STORAGE_DEFAULT_BUCKETS=$(RAW_AUDIO_STORAGE_DEFAULT_BUCKETS);
endef

# Processed audio storage service
define PROCESSED_AUDIO_STORAGE_ENVS
export PROCESSED_AUDIO_STORAGE_WEB_PORT=$(PROCESSED_AUDIO_STORAGE_WEB_PORT) \
PROCESSED_AUDIO_STORAGE_API_PORT=$(PROCESSED_AUDIO_STORAGE_API_PORT) \
PROCESSED_AUDIO_STORAGE_ROOT_USER=$(PROCESSED_AUDIO_STORAGE_ROOT_USER) \
PROCESSED_AUDIO_STORAGE_ROOT_PASSWORD=$(PROCESSED_AUDIO_STORAGE_ROOT_PASSWORD) \
PROCESSED_AUDIO_STORAGE_DEFAULT_BUCKETS=$(PROCESSED_AUDIO_STORAGE_DEFAULT_BUCKETS);
endef

# Packaged audio storage service
define PACKAGED_AUDIO_STORAGE_ENVS
export PACKAGED_AUDIO_STORAGE_WEB_PORT=$(PACKAGED_AUDIO_STORAGE_WEB_PORT) \
PACKAGED_AUDIO_STORAGE_API_PORT=$(PACKAGED_AUDIO_STORAGE_API_PORT) \
PACKAGED_AUDIO_STORAGE_ROOT_USER=$(PACKAGED_AUDIO_STORAGE_ROOT_USER) \
PACKAGED_AUDIO_STORAGE_ROOT_PASSWORD=$(PACKAGED_AUDIO_STORAGE_ROOT_PASSWORD) \
PACKAGED_AUDIO_STORAGE_DEFAULT_BUCKETS=$(PACKAGED_AUDIO_STORAGE_DEFAULT_BUCKETS);
endef

# Redpanda service
define REDPANDA_ENVS
export REDPANDA_KAFKA_INTERNAL_ADDR=$(REDPANDA_KAFKA_INTERNAL_ADDR) \
REDPANDA_KAFKA_EXTERNAL_ADDR=$(REDPANDA_KAFKA_EXTERNAL_ADDR) \
REDPANDA_KAFKA_ADVERTISE_INTERNAL_ADDR=$(REDPANDA_KAFKA_ADVERTISE_INTERNAL_ADDR) \
REDPANDA_KAFKA_ADVERTISE_EXTERNAL_ADDR=$(REDPANDA_KAFKA_ADVERTISE_EXTERNAL_ADDR) \
REDPANDA_PANDAPROXY_INTERNAL_ADDR=$(REDPANDA_PANDAPROXY_INTERNAL_ADDR) \
REDPANDA_PANDAPROXY_EXTERNAL_ADDR=$(REDPANDA_PANDAPROXY_EXTERNAL_ADDR) \
REDPANDA_PANDAPROXY_ADVERTISE_INTERNAL_ADDR=$(REDPANDA_PANDAPROXY_ADVERTISE_INTERNAL_ADDR) \
REDPANDA_PANDAPROXY_ADVERTISE_EXTERNAL_ADDR=$(REDPANDA_PANDAPROXY_ADVERTISE_EXTERNAL_ADDR) \
REDPANDA_SCHEMA_REGISTRY_INTERNAL_ADDR=$(REDPANDA_SCHEMA_REGISTRY_INTERNAL_ADDR) \
REDPANDA_SCHEMA_REGISTRY_EXTERNAL_ADDR=$(REDPANDA_SCHEMA_REGISTRY_EXTERNAL_ADDR) \
REDPANDA_RPC_ADDR=$(REDPANDA_RPC_ADDR) \
REDPANDA_ADVERTISE_RPC_ADDR=$(REDPANDA_ADVERTISE_RPC_ADDR) \
REDPANDA_MODE=$(REDPANDA_MODE) \
REDPANDA_SMP=$(REDPANDA_SMP) \
REDPANDA_LOG_LEVEL=$(REDPANDA_LOG_LEVEL) \
REDPANDA_IMAGE=$(REDPANDA_IMAGE) \
REDPANDA_DATA_DIR=$(REDPANDA_DATA_DIR) \
REDPANDA_SCHEMA_REGISTRY_PORT=$(REDPANDA_SCHEMA_REGISTRY_PORT) \
REDPANDA_PANDAPROXY_PORT=$(REDPANDA_PANDAPROXY_PORT) \
REDPANDA_KAFKA_PORT=$(REDPANDA_KAFKA_PORT) \
REDPANDA_ADMIN_PORT=$(REDPANDA_ADMIN_PORT) \
REDPANDA_INTERNAL_ADMIN_PORT=$(REDPANDA_INTERNAL_ADMIN_PORT);
endef

# Redpanda console service
define REDPANDA_CONSOLE_ENVS
export CONSOLE_IMAGE=$(CONSOLE_IMAGE) \
CONSOLE_CONFIG_FILEPATH=$(CONSOLE_CONFIG_FILEPATH) \
CONSOLE_PORT=$(CONSOLE_PORT);
endef

# Encapsulate environment variables and Docker Compose command
define DOCKER_COMPOSE_CMD
$(CATALOG_SERVICE_ENVS) \
$(CATALOG_STORAGE_ENVS) \
$(CATALOG_STORAGE_ADMIN_ENVS) \
$(SONG_STORAGE_ENVS) \
$(RAW_AUDIO_STORAGE_ENVS) \
$(PROCESSED_AUDIO_STORAGE_ENVS) \
$(PACKAGED_AUDIO_STORAGE_ENVS) \
$(REDPANDA_ENVS) \
$(REDPANDA_CONSOLE_ENVS) \
docker compose -f $(docker_compose_file)
endef

generate-acts:
	@echo "Generating acts..."
	@fish -c '$(deno_command) $(scripts_directory)/generate-acts.ts'
	@echo "Acts have been generated."

# Catalog service
.PHONY: catalog-service
catalog-service:
	@echo "Starting catalog service..."
	@$(CATALOG_SERVICE_ENVS) go run $(cmd_directory)/catalog/main.go
	@echo "Catalog service is up and running."

# Command to download songs
.PHONY: download-songs
download-songs:
	@echo "Starting song download script..."
	@SONGS_DIR=$(songs_directory) \
	$(python_command) $(scripts_directory)/download_music.py
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
	@read -p "Are you sure you want to rebuild all services without cache? [y/N] " confirm && [ $${confirm} = y ]
	@echo "Stopping services..."
	@$(DOCKER_COMPOSE_CMD) down --remove-orphans
	@echo "Deleting all cached files..."
	@sudo rm -rf $(deployments_directory)/data
	@echo "Rebuilding services without cache..."
	@$(DOCKER_COMPOSE_CMD) build --pull --no-cache
	@echo "Raising up services"
	@$(MAKE) start-services
	@echo "Services have been rebuilt and are running."

# Command to stop a single or multiple services
.PHONY: stop-services
stop-services:
	@read -p "Are you sure you want to stop these services? [y/N] " confirm && [ $${confirm} = y ]
	@echo "Stopping services..."
	@$(DOCKER_COMPOSE_CMD) down $(services)

# Command to stop all services
.PHONY: stop-all-services
stop-all-services:
	@read -p "Are you sure you want to stop all services? [y/N] " confirm && [ $${confirm} = y ]
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