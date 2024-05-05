FROM golang:1.16

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /api-go

EXPOSE 8080
CMD [ "/api-go" ]
