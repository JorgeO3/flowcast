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
python_command := "python"

# *** Authentication ***
# Authentication service 
auth_app_name := "auth-service"
auth_http_host := "0.0.0.0"
auth_http_port := "4100"
auth_db_url := \
  "postgresql://$(auth_postgres_user):$(auth_postgres_password)" \
  "@$(auth_postgres_host):5432/$(auth_postgres_db_name)?sslmode=disable"auth_version := "v1.0.0"
auth_log_level := "debug"
auth_migration_dir := $(migrations_directory)/auth

auth_account_email := jorge.testing9@gmail.com
auth_account_password := Jorgetesting1234

smtp_host := localhost
smtp_port := 1025

define AUTH_ENVS
export APP_NAME=$(auth_app_name) \
HTTP_HOST=$(auth_http_host) \
HTTP_PORT=$(auth_http_port) \
DB_NAME=$(auth_postgres_db_name) \
PG_URL=$(auth_database_url) \
VERSION=$(auth_version) \
LOG_LEVEL=$(auth_log_level) \
MIGRATIONS_PATH=$(auth_migration_dir) \
ACC_EMAIL=$(account_email) \
ACC_PASSWORD=$(account_password) \
SMTP_HOST=$(smtp_host) \
SMTP_PORT=$(smtp_port) \
EMAIL_TEMPLATE=$(email_template_file);
endef

# Authentication storage service
auth_storage_host := "localhost"
auth_storage_port := 5432
auth_storage_user := "jorge123"
auth_storage_password := "jorge123"
auth_storage_db_name := "auth_service"

define AUTH_STORAGE_ENVS
export AUTH_STORAGE_PORT=$(auth_postgres_port) \
AUTH_STORAGE_USER=$(auth_postgres_user) \
AUTH_STORAGE_PASSWORD=$(auth_postgres_password) \
AUTH_STORAGE_DB=$(auth_postgres_db_name);
endef

# Authentication storage admin service
auth_storage_admin_port := 8080

define AUTH_STORAGE_ADMIN_ENVS
export AUTH_STORAGE_ADMIN_PORT=$(auth_storage_admin_port);
endef

# Mail service
mailhog_smtp_port := "1025"
mailhog_http_port := "8025"

define MAILHOG_ENVS
export MAILHOG_SMTP_PORT=$(mailhog_service_smtp_port) \
MAILHOG_HTTP_PORT=$(mailhog_service_http_port); 
endef

# *** Catalog ***
# Catalog service

# Catalog storage service
catalog_storage_port := 27017
catalog_storage_root_user := "root"
catalog_storage_root_password := "root"
catalog_storage_server = "catalog-storage-service"

define CATALOG_STORAGE_ENVS
export CATALOG_STORAGE_PORT=$(catalog_storage_port) \
CATALOG_STORAGE_ROOT_USER=$(catalog_storage_root_user) \
CATALOG_STORAGE_ROOT_PASSWORD=$(catalog_storage_root_password);
endef

# Catalog storage admin service
catalog_storage_admin_port := 8081
catalog_storage_admin_user := $(catalog_storage_root_user)
catalog_storage_admin_password := $(catalog_storage_root_password)
catalog_storage_admin_url := \
  "mongodb://$(catalog_storage_root_user):$(catalog_storage_root_password)" \
  "@$(catalog_storage_server):27017/"
catalog_storage_admin_basicauth := false

define CATALOG_STORAGE_ADMIN_ENVS
export CATALOG_STORAGE_ADMIN_PORT=$(catalog_storage_admin_port) \
CATALOG_STORAGE_ADMIN_USER=$(catalog_storage_admin_user) \
CATALOG_STORAGE_ADMIN_PASSWORD=$(catalog_storage_admin_password) \
CATALOG_STORAGE_ADMIN_URL=$(catalog_storage_admin_url) \
CATALOG_STORAGE_ADMIN_BASICAUTH=$(catalog_storage_admin_basicauth);
endef

# *** Song ***
# Song service

# Song storage service configuration
song_storage_web_port := 9000
song_storage_api_port := 9001
song_storage_root_user := jorge123
song_storage_root_password := jorge123
song_storage_buckets := music-uploads,music-processed

define SONG_STORAGE_ENVS
export SONG_STORAGE_WEB_PORT=$(song_storage_web_port) \
SONG_STORAGE_API_PORT=$(song_storage_api_port) \
SONG_STORAGE_ROOT_USER=$(song_storage_root_user) \
SONG_STORAGE_ROOT_PASSWORD=$(song_storage_root_password) \
SONG_STORAGE_DEFAULT_BUCKETS=$(song_storage_buckets);
endef

# Encapsulate environment variables and Docker Compose command
define DOCKER_COMPOSE_CMD
$(AUTH_ENVS) \
$(AUTH_STORAGE_ENVS) \
$(AUTH_STORAGE_ADMIN_ENVS) \
$(MAILHOG_ENVS) \
$(CATALOG_STORAGE_ENVS) \
$(CATALOG_STORAGE_ADMIN_ENVS) \
$(SONG_STORAGE_ENVS) \
docker compose -f $(docker_compose_file)
endef

.PHONY: registration-request
registration-request:
	@echo "Starting auth registration request"
	http POST localhost:4100/register  \
	username=Jorge \
	email=jorge.testing9@gmail.com \
	password=Jorgetesting1234 \
	fullname="Jorge Osorio" \
	birthdate="1975-08-19T23:15:30.000Z" \
	gender=male \
	phone=323423423423 \
	emailNotif:=true \
	smsNotif:=true \
	socialLinks:='[{"platform": "facebook", "url": "facebook.com"}, {"platform": "github", "url": "github.com"}]'

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
	@$(DOCKER_COMPOSE_CMD) down
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
