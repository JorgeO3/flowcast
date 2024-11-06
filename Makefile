# Main Makefile

# Load global environment variables from .env if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Define root path and other directory paths
export project_root := $(CURDIR)
export go_root_path := gitlab.com/JorgeO3/flowcast

export songs_directory := $(project_root)/songs
export scripts_directory := $(project_root)/scripts
export deployments_directory := $(project_root)/deployments
export frontend_directory := $(project_root)/frontend
export backend_directory := $(project_root)/backend
export cmd_directory := $(project_root)/cmd
export migrations_directory := $(project_root)/migrations
export assets_directory := $(project_root)/assets

# Common commands
deno_command := deno run -A

# Include child Makefiles for different components
include deployments/Makefile
include web/Makefile       # Frontend Makefile
include cmd/Makefile       # Backend Makefile

# Declare phony targets to avoid conflicts with files of the same name
.PHONY: all generate-acts catalog-service download-songs

# Default target that builds and starts all necessary services
all: generate-acts catalog-service download-songs start-services start-frontend start-backend

# Generate acts using the specified Deno script
generate-acts:
	@echo "Generating acts..."
	@fish -c '$(deno_command) $(scripts_directory)/generate-acts.ts'
	@echo "Acts have been generated."

# Download songs by running the download script
download-songs:
	@echo "Starting the song download script..."
	@SONGS_DIR=$(songs_directory) \
	bash $(scripts_directory)/download_music.sh
	@echo "Song download completed."
