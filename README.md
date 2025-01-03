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

##http://localhost:4000/system-info

##docker build -t system-info-app .
##docker run -d -p 4000:4000 roxnatera/system-info-app:latest
##http://localhost:4000/system-info
