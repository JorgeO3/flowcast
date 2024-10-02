#!/bin/bash

set -e

# Función para mostrar mensajes de error
error_exit() {
  echo "ERROR: $1"
  exit 1
}

# Validar dependencias
if ! command -v docker &> /dev/null; then
  error_exit "Docker no está instalado. Instálalo y vuelve a intentarlo."
fi

# Verificar si el usuario puede ejecutar Docker
if ! docker info &> /dev/null; then
  error_exit "No se puede acceder al daemon de Docker. Asegúrate de que Docker está corriendo y tienes los permisos necesarios."
fi

# Variables
REGISTRY="harbor.local"  # Cambia esto según tu configuración
NAMESPACE="default"

# Lista de servicios
SERVICES=("service1" "service2" "api-service" "backend-service")

# Autenticarse en el registro Docker si es necesario
# Descomenta y ajusta las siguientes líneas si tu registro requiere autenticación
# echo "Iniciando sesión en el registro Docker..."
# docker login "${REGISTRY}" || error_exit "Falló el login en Docker Registry."

for service in "${SERVICES[@]}"; do
  echo "Procesando servicio: ${service}..."

  SERVICE_DIR="./services/${service}"
  
  if [ ! -d "${SERVICE_DIR}" ]; then
    echo "WARN: El directorio '${SERVICE_DIR}' no existe. Skipping..."
    continue
  fi

  IMAGE_TAG="${REGISTRY}/${NAMESPACE}/${service}:latest"

  echo "Construyendo imagen para ${service}..."
  docker build --pull -t "${IMAGE_TAG}" "${SERVICE_DIR}"

  echo "Empujando imagen para ${service}..."
  docker push "${IMAGE_TAG}"
done

echo "Todas las imágenes han sido construidas y empujadas exitosamente."
