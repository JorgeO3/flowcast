# deployments/Makefile

.PHONY: start-services rebuild-services stop-services stop-all-services view-logs cleanup-docker check-status

# Definir las rutas y archivos específicos de deployments
deployments_directory := $(CURDIR)
docker_compose_file := $(deployments_directory)/docker-compose.yaml
env_file := $(deployments_directory)/.env

# Cargar y exportar variables de entorno específicas para deployments
ifneq (,$(wildcard $(env_file)))
    include $(env_file)
    export
endif

# Comando Docker Compose encapsulado con variables de entorno
DOCKER_COMPOSE_CMD := docker compose -f $(docker_compose_file)

# Documentación de las recetas
# Iniciar servicios de deployments
start-services:
	@echo "Iniciando servicios de deployments..."
	@$(DOCKER_COMPOSE_CMD) up -d
	@echo "Servicios de deployments están activos."

# Reconstruir servicios sin cache
rebuild-services:
	@read -p "¿Estás seguro de que deseas reconstruir todos los servicios sin cache? [y/N] " confirm && [ "$${confirm}" = y ]
	@echo "Deteniendo servicios..."
	@$(DOCKER_COMPOSE_CMD) down --remove-orphans
	@echo "Eliminando archivos cacheados..."
	@sudo rm -rf $(deployments_directory)/data
	@echo "Reconstruyendo servicios sin cache..."
	@$(DOCKER_COMPOSE_CMD) build --pull --no-cache
	@echo "Reiniciando servicios..."
	@$(MAKE) -C $(deployments_directory) start-services
	@echo "Servicios han sido reconstruidos y están corriendo."

# Detener servicios específicos
stop-services:
	@read -p "¿Estás seguro de que deseas detener estos servicios? [y/N] " confirm && [ "$${confirm}" = y ]
	@echo "Deteniendo servicios específicos..."
	@$(DOCKER_COMPOSE_CMD) down $(services)

# Detener todos los servicios
stop-all-services:
	@read -p "¿Estás seguro de que deseas detener todos los servicios? [y/N] " confirm && [ "$${confirm}" = y ]
	@echo "Deteniendo todos los servicios..."
	@$(DOCKER_COMPOSE_CMD) down
	@echo "Todos los servicios han sido detenidos."

# Ver logs de todos los servicios
view-logs:
	@echo "Mostrando logs de todos los servicios..."
	@$(DOCKER_COMPOSE_CMD) logs -f

# Limpiar objetos Docker no usados
cleanup-docker:
	@echo "Eliminando todos los objetos Docker no usados..."
	@docker system prune -f
	@echo "Limpieza completada."

# Verificar estado de servicios
check-status:
	@echo "Verificando el estado de los servicios..."
	@$(DOCKER_COMPOSE_CMD) ps
	@echo "Verificación de estado completada."
