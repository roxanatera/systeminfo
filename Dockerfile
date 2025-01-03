# Usa una imagen base de Go
FROM golang:1.20 AS builder

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente de la aplicación
COPY . .

# Compila la aplicación
RUN go build -o system-info .

# Usa una imagen más ligera para ejecutar la aplicación
FROM alpine:latest

# Copia el binario compilado desde el builder
COPY --from=builder /app/system-info /system-info

# Expone el puerto 4000
EXPOSE 4000

# Comando por defecto para ejecutar la aplicación
CMD ["/system-info"]
