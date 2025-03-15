#!/bin/bash

# URL de tu API en AWS
API_URL="http://54.85.123.137/"

# Función para verificar el estado de la API
check_api_status() {
    # Intenta hacer una solicitud HTTP a la API
    response=$(curl -s -o /dev/null -w "%{http_code}" $API_URL)

    if [[ $response -eq 200 ]]; then
        echo "API is up and running"
        echo "status=success" >> $GITHUB_OUTPUT
    elif [[ $response -eq 000 ]]; then
        echo "API is down (off)"
        echo "status=off" >> $GITHUB_OUTPUT
    else
        echo "API returned an error (HTTP code: $response)"
        echo "status=failed" >> $GITHUB_OUTPUT
    fi
}

# Ejecutar la verificación
check_api_status