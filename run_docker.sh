#!/bin/bash

# Script para rodar a aplicação stress-test via Docker

# Verifica se os parâmetros foram passados
if [ -z "$1" ] || [ -z "$2" ] || [ -z "$3" ]; then
  echo "Uso: ./run_docker.sh <URL> <REQUESTS> <CONCURRENCY>"
  echo "Exemplo: ./run_docker.sh http://google.com 10 10"
  exit 1
fi

URL=$1
REQUESTS=$2
CONCURRENCY=$3

# Constrói a imagem Docker (se ainda não foi construída ou se o código mudou)
docker build -t stress-test-app .

# Executa o container Docker com os parâmetros fornecidos
docker run stress-test-app --url="$URL" --requests="$REQUESTS" --concurrency="$CONCURRENCY"
