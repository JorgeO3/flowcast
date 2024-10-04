# Makefile principal

# Cargar variables de entorno globales desde .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Definir las rutas raíz y otros directorios
project_root := $(CURDIR)
go_root_path := gitlab.com/JorgeO3/flowcast

songs_directory := $(project_root)/songs
scripts_directory := $(project_root)/scripts
deployments_directory := $(project_root)/deployments
frontend_directory := $(project_root)/frontend
backend_directory := $(project_root)/backend
cmd_directory := $(project_root)/cmd
migrations_directory := $(project_root)/migrations
assets_directory := $(project_root)/assets

# Comandos comunes
python_command := python
deno_command := deno run -A

# Incluir Makefiles hijos
include deployments/Makefile
include web/Makefile
include cmd/Makefile

# Documentación de las recetas
.PHONY: all generate-acts catalog-service download-songs

# Objetivo por defecto
all: generate-acts catalog-service download-songs start-services start-frontend start-backend

# Generar acts
generate-acts:
	@echo "Generando acts..."
	@fish -c '$(deno_command) $(scripts_directory)/generate-acts.ts'
	@echo "Acts han sido generados."

# Iniciar el servicio de catálogo
catalog-service:
	@echo "Iniciando el servicio de catálogo..."
	@$(CATALOG_SERVICE_ENVS) go run $(cmd_directory)/catalog/main.go
	@echo "Servicio de catálogo está activo."

# Descargar canciones
download-songs:
	@echo "Iniciando el script de descarga de canciones..."
	@SONGS_DIR=$(songs_directory) \
	$(python_command) $(scripts_directory)/download_music.py
	@echo "Descarga de canciones completada."
