#!/bin/bash

set -e

# Variables
REGISTRY="harbor.local"  # Cambia esto según tu configuración
NAMESPACE="default"

# Lista de servicios
SERVICES=("service1" "service2" "api-service" "backend-service")

for service in "${SERVICES[@]}"; do
  echo "Construyendo imagen para ${service}..."
  docker build -t ${REGISTRY}/${NAMESPACE}/${service}:latest ./services/${service}

  echo "Empujando imagen para ${service}..."
  docker push ${REGISTRY}/${NAMESPACE}/${service}:latest
done

echo "Todas las imágenes han sido construidas y empujadas exitosamente."
