# Información del Sistema Cliente y Servidor

Este proyecto es una aplicación desarrollada en **Go** que sirve una página web para mostrar información detallada del sistema operativo y navegador del cliente, así como información básica del sistema del servidor. La aplicación está diseñada para ejecutarse localmente o desplegarse en plataformas como **Render**.

## Funcionalidades

### Cliente
La aplicación recopila y muestra la siguiente información del cliente (navegador):
- **Sistema Operativo**: Detecta el sistema operativo del cliente.
- **Navegador**: Identifica el navegador utilizado.
- **Resolución de Pantalla**: Muestra el ancho y alto de la pantalla del cliente en píxeles.
- **Idioma del Navegador**: Detecta el idioma configurado en el navegador.
- **Zona Horaria**: Muestra la zona horaria del cliente.
- **Soporte de Cookies**: Indica si las cookies están habilitadas.
- **Información de la Conexión de Red**: Muestra el tipo de conexión y la velocidad si el navegador lo soporta.
- **Plataforma del Dispositivo**: Muestra información sobre la plataforma del dispositivo.
- **Soporte de APIs Específicas**:
  - WebRTC
  - Geolocalización
  - WebAssembly
- **Geolocalización**: Si el cliente lo permite, muestra las coordenadas de ubicación (latitud y longitud).

### Servidor
La aplicación también incluye un endpoint `/systeminfo` que muestra información básica sobre el servidor:
- **Sistema Operativo del Servidor**.
- **Arquitectura del Servidor**.
- **Número de CPUs**.

### Navegación Sencilla
Incluye un enlace en la página principal para acceder a la información del sistema del servidor.

## Tecnologías Utilizadas

- **Backend**:
  - Lenguaje: Go
  - Paquetes estándar: `net/http`, `os`, `runtime`

- **Frontend**:
  - HTML, CSS y JavaScript
  - Uso de APIs modernas como `navigator.userAgent`, `navigator.geolocation`, y `Intl.DateTimeFormat`.

## Requisitos Previos

- Tener instalado **Go** (versión 1.20 o superior recomendada).
- Un navegador moderno que soporte las APIs utilizadas.

## Estructura del Proyecto

