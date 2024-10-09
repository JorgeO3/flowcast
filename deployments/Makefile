# deployments/Makefile

.PHONY: start-services rebuild-services stop-services stop-all-services view-logs cleanup-docker check-status

# Define paths and files specific to deployments
deployments_directory := $(CURDIR)
docker_compose_file := $(deployments_directory)/docker-compose.yaml
env_file := $(deployments_directory)/.env

# Load and export environment variables specific to deployments if .env exists
ifneq (,$(wildcard $(env_file)))
    include $(env_file)
    export
endif

# Docker Compose command encapsulated with the specified docker-compose file
DOCKER_COMPOSE_CMD := docker compose -f $(docker_compose_file)

# Start deployment services
start-services:
	@echo "Starting deployment services..."
	@$(DOCKER_COMPOSE_CMD) up -d
	@echo "Deployment services are now active."

# Rebuild services without using the cache
rebuild-services:
	@read -p "Are you sure you want to rebuild all services without cache? [y/N] " confirm && [ "$${confirm}" = y ]
	@echo "Stopping services..."
	@$(DOCKER_COMPOSE_CMD) down --remove-orphans
	@echo "Removing cached files..."
	@sudo rm -rf $(deployments_directory)/data
	@echo "Rebuilding services without cache..."
	@$(DOCKER_COMPOSE_CMD) build --pull --no-cache
	@echo "Restarting services..."
	@$(MAKE) -C $(deployments_directory) start-services
	@echo "Services have been rebuilt and are running."

# Stop specific services
stop-services:
	@read -p "Are you sure you want to stop these services? [y/N] " confirm && [ "$${confirm}" = y ]
	@echo "Stopping specified services..."
	@$(DOCKER_COMPOSE_CMD) down $(services)

# Stop all services
stop-all-services:
	@read -p "Are you sure you want to stop all services? [y/N] " confirm && [ "$${confirm}" = y ]
	@echo "Stopping all services..."
	@$(DOCKER_COMPOSE_CMD) down
	@echo "All services have been stopped."

# View logs of all services
view-logs:
	@echo "Displaying logs for all services..."
	@$(DOCKER_COMPOSE_CMD) logs -f

# Clean up unused Docker objects
cleanup-docker:
	@echo "Removing all unused Docker objects..."
	@docker system prune -f
	@echo "Cleanup completed."

# Check the status of services
check-status:
	@echo "Checking the status of services..."
	@$(DOCKER_COMPOSE_CMD) ps
	@echo "Status check completed."
