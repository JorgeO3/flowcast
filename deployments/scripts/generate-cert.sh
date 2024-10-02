#!/bin/bash

set -e

# Validar que OpenSSL está instalado
if ! command -v openssl &> /dev/null; then
  echo "ERROR: OpenSSL no está instalado. Instálalo y vuelve a intentarlo."
  exit 1
fi

# Validar que kubectl está instalado
if ! command -v kubectl &> /dev/null; then
  echo "ERROR: kubectl no está instalado o no está en el PATH."
  exit 1
fi

# Validar que el cluster está activo
if ! kubectl cluster-info &> /dev/null; then
  echo "ERROR: No se pudo acceder al cluster de Kubernetes. Asegúrate de que está corriendo."
  exit 1
fi

# Directorio donde se almacenarán los certificados generados
CERT_DIR="$(dirname "$0")/../certs"

# Array de servicios con sus dominios y namespaces
declare -A SERVICES=(
  ["nginx"]="api-gateway.local:default"
  ["varnish"]="varnish.local:varnish"
  ["grafana"]="grafana.local:monitoring"
  ["harbor"]="harbor.local:default"
)

# Función para crear namespaces si no existen
create_namespace() {
  local NAMESPACE=$1
  if ! kubectl get namespace "${NAMESPACE}" &> /dev/null; then
    echo "Creando namespace '${NAMESPACE}'..."
    kubectl create namespace "${NAMESPACE}"
  else
    echo "Namespace '${NAMESPACE}' ya existe."
  fi
}

# Función para generar certificados
generate_cert() {
  local SERVICE_NAME=$1
  local DOMAIN=$2
  local NAMESPACE=$3

  local KEY_FILE="${CERT_DIR}/${SERVICE_NAME}.key"
  local CRT_FILE="${CERT_DIR}/${SERVICE_NAME}.crt"
  local SUBJECT="/CN=${DOMAIN}/O=${SERVICE_NAME^}"

  # Generar la clave privada si no existe
  if [ ! -f "${KEY_FILE}" ]; then
    echo "Generando clave privada para ${SERVICE_NAME}..."
    openssl genrsa -out "${KEY_FILE}" 2048
  else
    echo "Clave privada para ${SERVICE_NAME} ya existe."
  fi

  # Generar el certificado si no existe
  if [ ! -f "${CRT_FILE}" ]; then
    echo "Generando certificado TLS para ${SERVICE_NAME} (${DOMAIN})..."
    openssl req -x509 -nodes -days 365 -new -key "${KEY_FILE}" \
      -out "${CRT_FILE}" \
      -subj "${SUBJECT}"
  else
    echo "Certificado TLS para ${SERVICE_NAME} ya existe."
  fi

  # Crear o actualizar el Secret de Kubernetes
  echo "Creando/Actualizando Secret TLS para ${SERVICE_NAME} en el namespace '${NAMESPACE}'..."
  kubectl create secret tls "${SERVICE_NAME}-tls" \
    --cert="${CRT_FILE}" \
    --key="${KEY_FILE}" \
    --namespace="${NAMESPACE}" \
    --dry-run=client -o yaml | kubectl apply -f -
}

# Crear el directorio de certificados si no existe
mkdir -p "${CERT_DIR}"

# Iterar sobre los servicios y generar certificados
for SERVICE in "${!SERVICES[@]}"; do
  IFS=':' read -r DOMAIN NAMESPACE <<< "${SERVICES[$SERVICE]}"
  create_namespace "${NAMESPACE}"
  generate_cert "${SERVICE}" "${DOMAIN}" "${NAMESPACE}"
done

echo "Todos los certificados y Secrets han sido generados y aplicados exitosamente."
