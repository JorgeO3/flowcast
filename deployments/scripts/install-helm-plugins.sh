#!/bin/bash

# Lista de plugins a instalar
PLUGINS=("diff")

for plugin in "${PLUGINS[@]}"; do
  if ! helm plugin list | grep -q "^${plugin}\b"; then
    echo "Instalando el plugin de Helm: ${plugin}"
    helm plugin install https://github.com/databus23/helm-diff
    if [ $? -ne 0 ]; then
      echo "Error: Falló la instalación del plugin ${plugin}"
      exit 1
    fi
  else
    echo "El plugin de Helm '${plugin}' ya está instalado."
  fi
done

exit 0
