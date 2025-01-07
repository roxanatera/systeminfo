package main

import (
    "fmt"
    "net/http"
    "os"
    "runtime"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
)

func systemInfo(w http.ResponseWriter, r *http.Request) {
    // Establece el tipo de contenido y la codificación de caracteres
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    // Comienza a generar la página HTML
    fmt.Fprintf(w, "<h1>Información del Sistema Operativo</h1>")
    fmt.Fprintf(w, "<strong>Sistema Operativo:</strong> %s<br>", runtime.GOOS)
    fmt.Fprintf(w, "<strong>Arquitectura:</strong> %s<br>", runtime.GOARCH)
    fmt.Fprintf(w, "<strong>Número de CPUs:</strong> %d<br>", runtime.NumCPU())

    // Información de disco
    d, _ := disk.Usage("/")
    fmt.Fprintf(w, "<strong>Espacio total del disco:</strong> %v GB<br>", d.Total/1024/1024/1024)
    fmt.Fprintf(w, "<strong>Espacio usado del disco:</strong> %v GB (%.2f%%)<br>", d.Used/1024/1024/1024, d.UsedPercent)

    // Información de la memoria RAM
    v, _ := mem.VirtualMemory()
    fmt.Fprintf(w, "<strong>Total memoria RAM:</strong> %v GB<br>", v.Total/1024/1024/1024)
    fmt.Fprintf(w, "<strong>Memoria RAM usada:</strong> %v GB (%.2f%%)<br>", v.Used/1024/1024/1024, v.UsedPercent)

    // Directorio del usuario
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Fprintf(w, "<strong>Error obteniendo el directorio del usuario:</strong> %s<br>", err)
    } else {
        fmt.Fprintf(w, "<strong>Carpeta del Usuario:</strong> %s<br>", homeDir)
    }
}

func main() {
    http.HandleFunc("/", systemInfo) // Establece el handler para la raíz
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Puerto por defecto si no se especifica
    }
    fmt.Printf("Servidor iniciado en http://localhost:%s\n", port)
    http.ListenAndServe(":"+port, nil)
}
