# backend/Makefile

.PHONY: start-backend build-backend clean-backend

# Definir rutas y archivos específicos del backend
env_file := $(cmd_directory)/.env

# Cargar y exportar variables de entorno específicas para backend
ifneq (,$(wildcard $(env_file)))
    include $(env_file)
    export
endif

# Start the catalog service
catalog-service:
	@echo "Starting the catalog service..."
#	@go run -tags grpc $(cmd_directory)/catalog/main.go
	@go run $(cmd_directory)/catalog/main.go
	@echo "Catalog service is now active."

# Start the transcoding service
transcode-service:
#	@go run -tags grpc $(cmd_directory)/transcode/main.go
	@go run $(cmd_directory)/transcode/main.go
	@echo "Transcoding service is now active."