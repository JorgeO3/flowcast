# backend/Makefile

.PHONY: start-backend build-backend clean-backend

# Definir rutas y archivos específicos del backend
backend_directory := $(CURDIR)
env_file := $(backend_directory)/.env

# Cargar y exportar variables de entorno específicas para backend
ifneq (,$(wildcard $(env_file)))
    include $(env_file)
    export
endif

# Iniciar el backend
start-backend:
	@echo "Iniciando el backend..."
	@cd $(backend_directory) && go build -o backend main.go
	@cd $(backend_directory) && ./backend
	@echo "Backend está activo."

# Construir el backend para producción
build-backend:
	@echo "Construyendo el backend para producción..."
	@cd $(backend_directory) && go build -o backend main.go
	@echo "Construcción del backend completada."

# Limpiar archivos generados del backend
clean-backend:
	@echo "Limpiando archivos del backend..."
	@cd $(backend_directory) && go clean
	@echo "Limpieza del backend completada."
