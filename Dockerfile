FROM golang:1.23.4-alpine

# Instalar dependencias y wait-for-it
RUN apk add --no-cache bash curl mysql-client && \
    curl -o /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x /usr/local/bin/wait-for-it.sh

# Variables de entorno para Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Copiar archivos necesarios para descargar dependencias
COPY .env ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar todo el código fuente
COPY . .

# Construir la aplicación
RUN go build -o main .

CMD ["sh", "-c", "/usr/local/bin/wait-for-it.sh db:3306 --timeout=30 -- ./main"]