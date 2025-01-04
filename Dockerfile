# Usa una imagen base de Go para construir el binario
FROM golang:1.23 AS builder

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia los archivos necesarios
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Compila la aplicación
RUN go build -o system-info .

# Usa una imagen base para ejecutar el binario
FROM alpine:latest

# Instala las dependencias necesarias
RUN apk add --no-cache libc6-compat

# Copia el binario compilado desde el builder
COPY --from=builder /app/system-info /system-info


RUN chmod +x /system-info

# Expone el puerto 4000
EXPOSE 4000

# Expone el puerto (opcional, depende de si quieres fijarlo)
EXPOSE $PORT

# Comando por defecto para ejecutar la aplicación
CMD ["/system-info"]
