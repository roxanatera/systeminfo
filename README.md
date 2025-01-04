# System Info App

System Info App es una aplicación desarrollada en **Go (Golang)** que recopila y muestra información del sistema, como CPU, memoria, disco y más. También incluye una API que expone esta información en formato JSON.

## Características

- Obtención de información detallada del sistema:
  - Sistema operativo, hostname y tiempo de actividad.
  - Uso de CPU, memoria y disco.
- Exposición de datos a través de una API HTTP en formato JSON.
- Generación de una salida HTML estilizada para visualizar los datos.
- Desplegable en **Docker** para facilitar su uso.

## Tecnologías utilizadas

- **Backend:** Golang
- **Framework HTTP:** Fiber
- **Biblioteca de Sistema:** gopsutil
- **Contenedores:** Docker

## Cómo usar

### Ejecución local

1. Clona este repositorio:
   ```bash
   git clone  https://github.com/roxanatera/systeminfo.git
git branch -M main

###go build -o system-info .
./system-info
## Para descargar la imagen, ejecuta:
**docker pull roxanatera/system-info-app:latest**


## http://localhost:4000/system-info

## docker build -t system-info-app .
**Para ejecutar la imagen de Docker y obtener información del sistema anfitrión, usa el siguiente comando:**
## docker run -d -p 4000:4000 --privileged -v /:/host roxanatera/system-info-app:latest

## http://localhost:4000/system-info

## Ejemplo de Salida JSON

Al consultar el endpoint `/system-info`, se obtiene una respuesta en formato JSON con la siguiente estructura:

```json
{
    "host_info": {
        "os": "linux",
        "platform": "ubuntu",
        "version": "20.04",
        "hostname": "my-host",
        "uptime": 123456
    },
    "cpu_info": {
        "model_name": "Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz",
        "cores": 6,
        "usage": 15.7
    },
    "memory_info": {
        "total_gb": 16.0,
        "used_gb": 8.5,
        "usage_percent": 53.1
    },
    "disk_info": {
        "total_gb": 512.0,
        "used_gb": 256.0,
        "usage_percent": 50.0
    }
}
