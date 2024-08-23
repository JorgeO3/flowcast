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

# Song service
# Song storage service configuration

# Encapsulate environment variables and Docker Compose command
define DOCKER_COMPOSE_CMD
$(CATALOG_SERVICE_ENVS) \
$(CATALOG_STORAGE_ENVS) \
$(CATALOG_STORAGE_ADMIN_ENVS) \
$(SONG_STORAGE_ENVS) \
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