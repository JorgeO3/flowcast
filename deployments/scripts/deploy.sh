#!/bin/bash

set -e

# Función para mostrar mensajes de error
error_exit() {
  echo "ERROR: $1"
  exit 1
}

# Verificar dependencias
for cmd in kind kubectl helmfile; do
  if ! command -v "$cmd" &> /dev/null; then
    error_exit "$cmd no está instalado. Por favor, instálalo y vuelve a intentarlo."
  fi
done

# Variables
CLUSTER_NAME="dev-cluster"
HELMFILE_PATH="$(dirname "$0")/../helmfile.yaml"
KIND_CONFIG_PATH="$(dirname "$0")/../kind-config.yaml"

# Verificar si kind-config.yaml existe
if [ ! -f "${KIND_CONFIG_PATH}" ]; then
  error_exit "No se encontró el archivo kind-config.yaml en ${KIND_CONFIG_PATH}"
fi

# Crear el cluster si no existe
if ! kind get clusters | grep -q "^${CLUSTER_NAME}$"; then
  echo "Creando el clúster de Kind..."
  kind create cluster --name "${CLUSTER_NAME}" --config "${KIND_CONFIG_PATH}"
else
  echo "El clúster de Kind '${CLUSTER_NAME}' ya existe."
fi

# Configurar kubectl
kubectl config use-context "kind-${CLUSTER_NAME}"
if [ $? -ne 0 ]; then
  error_exit "No se pudo cambiar al contexto 'kind-${CLUSTER_NAME}'."
fi

# Instalar operadores si existen
echo "Instalando operadores de Kubernetes..."
if ls operators/*/ >/dev/null 2>&1; then
  for operator in operators/*/; do
    echo "Desplegando operador: ${operator}"
    if ls "${operator}"/*.yaml >/dev/null 2>&1; then
      kubectl apply -f "${operator}"
    else
      echo "WARN: No se encontraron archivos YAML en ${operator}. Skipping..."
    fi
  done
else
  echo "No se encontraron operadores para desplegar."
fi

# Manejo de errores para helmfile
trap 'echo "ERROR: Falló el despliegue con Helmfile."; exit 1' ERR

# Desplegar con Helmfile
echo "Desplegando servicios con Helmfile..."
helmfile -f "${HELMFILE_PATH}" apply

echo "Despliegue completado exitosamente."
