#!/bin/bash

# Defina o nome da aplicação
APP_NAME="user-register"
PROJECT_ROOT="."


# Executa go mod tidy
go mod tidy

#removendo o app anterior
sudo rm -rf  ${PROJECT_ROOT}/_deploy/

# Compila o projeto e define o nome do executável como "mycli"
go build -ldflags="-w -s" -o ${PROJECT_ROOT}/_deploy/${APP_NAME} ${PROJECT_ROOT}/cmd/app/main.go

# Navega para o diretório build
cd ${PROJECT_ROOT}/build

# Executa o aplicativo compilado
./${APP_NAME}