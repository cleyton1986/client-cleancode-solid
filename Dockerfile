# Argumento para definir o nome da aplicação
ARG APP_NAME=server-api-go

# Usando a imagem do Go como builder
FROM golang:latest as builder
ARG APP_NAME

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando todo o projeto para o container
COPY . .

# Compilando o projeto
RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-w -s" -o bin/${APP_NAME} cmd/app/main.go

# Usando a imagem do Go para debug
FROM golang:latest AS debugger
WORKDIR /app
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . .

# Usando Alpine para as dependências
FROM alpine:3.14 as dependencies
WORKDIR /app
ENV ORACLE_HOME=/opt/oracle/instantclient
RUN mkdir -p $ORACLE_HOME

RUN apk add --no-cache libc6-compat

# Instalando as dependências necessárias no Alpine
RUN apk update && \
  apk add --no-cache libssl1.1 curl unzip bash build-base && \
  curl -o /tmp/instantclient-basic-linux.x64-21.1.0.0.0.zip https://download.oracle.com/otn_software/linux/instantclient/211000/instantclient-basic-linux.x64-21.1.0.0.0.zip && \
  curl -o /tmp/instantclient-sdk-linux.x64-21.1.0.0.0.zip https://download.oracle.com/otn_software/linux/instantclient/211000/instantclient-sdk-linux.x64-21.1.0.0.0.zip && \
  cd /tmp && \
  unzip instantclient-basic-linux.x64-21.1.0.0.0.zip && \
  unzip instantclient-sdk-linux.x64-21.1.0.0.0.zip && \
  mv /tmp/instantclient_21_1/* $ORACLE_HOME && \
  rm -rf /tmp/instantclient_21_1 /var/cache/apk/*


# Configurando a imagem final
FROM dependencies AS app
ARG APP_NAME
ENV APP_NAME=${APP_NAME}
ENV ORACLE_HOME=/opt/oracle/instantclient
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$ORACLE_HOME
ENV TNS_ADMIN=$ORACLE_HOME/network/admin

WORKDIR /app

# Adicionando o binário do aplicativo
COPY --from=builder /app/bin/${APP_NAME} ./app/

# Adicionando o arquivo .env
COPY .env /app/

# Definindo a variável de ambiente PATH
ENV PATH=$PATH:$ORACLE_HOME
EXPOSE 8000
WORKDIR /app/
ENTRYPOINT ./app/${APP_NAME}

