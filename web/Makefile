# frontend/Makefile

.PHONY: start-frontend build-frontend clean-frontend

# Definir rutas y archivos específicos del frontend
frontend_directory := $(CURDIR)
env_file := $(frontend_directory)/.env

# Cargar y exportar variables de entorno específicas para frontend
ifneq (,$(wildcard $(env_file)))
    include $(env_file)
    export
endif

# Iniciar el frontend
start-frontend:
	@echo "Iniciando el frontend..."
	@cd $(frontend_directory) && npm install
	@cd $(frontend_directory) && npm start
	@echo "Frontend está activo."

# Construir el frontend para producción
build-frontend:
	@echo "Construyendo el frontend para producción..."
	@cd $(frontend_directory) && npm run build
	@echo "Construcción del frontend completada."

# Limpiar archivos generados del frontend
clean-frontend:
	@echo "Limpiando archivos del frontend..."
	@cd $(frontend_directory) && npm run clean
	@echo "Limpieza del frontend completada."
