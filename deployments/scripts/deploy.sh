#!/bin/bash

set -e

# Variables
CLUSTER_NAME="dev-cluster"
HELMFILE_PATH="$(dirname "$0")/../helmfile.yaml"
KIND_CONFIG_PATH="$(dirname "$0")/../kind-config.yaml"

# Verificar si kind-config.yaml existe
if [ ! -f "${KIND_CONFIG_PATH}" ]; then
  echo "ERROR: No se encontró el archivo kind-config.yaml en ${KIND_CONFIG_PATH}"
  exit 1
fi

# Crear el cluster si no existe
if ! kind get clusters | grep -q "^${CLUSTER_NAME}$"; then
  echo "Creando el clúster de Kind..."
  kind create cluster --name ${CLUSTER_NAME} --config ${KIND_CONFIG_PATH}
else
  echo "El clúster de Kind '${CLUSTER_NAME}' ya existe."
fi

# Configurar kubectl
kubectl config use-context kind-${CLUSTER_NAME}

# Instalar operadores si existen
echo "Instalando operadores de Kubernetes..."
if ls operators/*/ >/dev/null 2>&1; then
  for operator in operators/*/; do
    echo "Desplegando operador: ${operator}"
    kubectl apply -f "${operator}"
  done
else
  echo "No se encontraron operadores para desplegar."
fi

# Desplegar con Helmfile
echo "Desplegando servicios con Helmfile..."
helmfile -f ${HELMFILE_PATH} apply

echo "Despliegue completado exitosamente."
