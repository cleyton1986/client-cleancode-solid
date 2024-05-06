#!/bin/bash

#!/bin/bash
echo "======================================================================================"
echo ""
echo "██████╗ ██╗   ██╗██╗██╗     ██████╗      █████╗ ██████╗ ██████╗      ██████╗  ██████╗ "
echo "██╔══██╗██║   ██║██║██║     ██╔══██╗    ██╔══██╗██╔══██╗██╔══██╗    ██╔════╝ ██╔═══██╗"
echo "██████╔╝██║   ██║██║██║     ██║  ██║    ███████║██████╔╝██████╔╝    ██║  ███╗██║   ██║"
echo "██╔══██╗██║   ██║██║██║     ██║  ██║    ██╔══██║██╔═══╝ ██╔═══╝     ██║   ██║██║   ██║"
echo "██████╔╝╚██████╔╝██║███████╗██████╔╝    ██║  ██║██║     ██║         ╚██████╔╝╚██████╔╝"
echo "╚═════╝  ╚═════╝ ╚═╝╚══════╝╚═════╝     ╚═╝  ╚═╝╚═╝     ╚═╝          ╚═════╝  ╚═════╝ "
echo ""
echo "======================================================================================"



# Defina o nome da aplicação
APP_NAME="server-api-go"

ROOT_DIR="."

# Diretório do projeto onde o arquivo main.go e go.mod estão localizados
PROJECT_DIR="${ROOT_DIR}/cmd/app"

# Diretório para o binário compilado, na raiz do projeto, no mesmo nível que "cmd"
BIN_DIR="${ROOT_DIR}/bin"

# Defina o nome da imagem
IMAGE_NAME="${APP_NAME}-image"

# Defina o nome do container
CONTAINER_NAME="${APP_NAME}-container"
echo "======================================================================================"
# Atualiza as dependências do Go
echo "Atualizando as dependências..."
cd $PROJECT_DIR
go mod tidy
if [ $? -ne 0 ]; then
    echo "Falha ao atualizar as dependências."
    exit 1
fi

echo "======================================================================================"
# Remove a versão anterior do binário
echo "Removendo o binário anterior..."
sudo rm -rf ../../$BIN_DIR

echo "======================================================================================"
# Fazendo o build do binário Go
echo "Construindo o binário Go..."
go build -o ../../$BIN_DIR/${APP_NAME} main.go
if [ $? -ne 0 ]; then
    echo "Falha ao construir o binário Go."
    exit 1

else
    cd -
    echo "Construindo com sucesso!!"
fi

echo "======================================================================================"
# Verificar se o container existe e pará-lo se existir
if [ $(docker ps -a -q -f name=^/${CONTAINER_NAME}$) ]; then
    echo "Parando e removendo o container existente..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
else
    echo "Nenhum container para parar ou remover."
fi

echo "======================================================================================"
# Construir a imagem Docker, passando a variável APP_NAME
echo "Construindo a imagem Docker..."
docker build --build-arg APP_NAME=$APP_NAME -t $IMAGE_NAME .
echo "======================================================================================"

# Rodar o container Docker
echo "Rodando o container Docker..."
docker run -d --name $CONTAINER_NAME -p 8000:8000 $IMAGE_NAME

echo "======================================================================================"
