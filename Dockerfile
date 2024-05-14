# Argumento para definir o nome da aplicação
ARG APP_NAME=server-api-go

# Usando a imagem do Go como builder
FROM golang:1.22 as builder
ARG APP_NAME

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando o arquivo go.mod e go.sum e baixando as dependências
COPY go.mod go.sum ./
RUN go mod vendor

COPY vendor/ ./
RUN go mod download

# Copiando o restante do projeto para o container
COPY . .

# Compilando o projeto
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/${APP_NAME} ./cmd/app/main.go

# Usando a imagem Distroless para a aplicação final
FROM gcr.io/distroless/base
ARG APP_NAME
COPY --from=builder /bin/${APP_NAME} /app/${APP_NAME}

# Definindo a porta e o comando de execução
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["/app/server-api-go"]  # Usar o nome explícito do binário

