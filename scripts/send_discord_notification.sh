#!/bin/bash

# Leer el estado del despliegue desde el argumento
status=$1

# URL del webhook de pruebas de Discord
DISCORD_WEBHOOK_URL=$DISCORD_WEBHOOK_TEST

# Mensaje predeterminado
message=""

# Definir el mensaje según el estado
if [[ $status == "success" ]]; then
  message="✅ **Despliegue Exitoso**\nEl API está funcionando correctamente."
elif [[ $status == "failed" ]]; then
  message="❌ **Despliegue Fallido**\nEl API está caída o devolvió un error."
elif [[ $status == "off" ]]; then
  message="⚠️ **API Apagada**\nEl API está apagada. No se requiere acción adicional."
else
  message="🔔 **Estado Desconocido**\nEl estado del despliegue no pudo ser determinado."
fi

# Enviar el mensaje a Discord
curl -X POST -H "Content-Type: application/json" \
  -d "{\"content\": \"$message\"}" \
  $DISCORD_WEBHOOK_URL