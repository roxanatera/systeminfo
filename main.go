package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"log"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func systemInfo(w http.ResponseWriter, r *http.Request) {
	// Establece el tipo de contenido y la codificación de caracteres
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Genera la página HTML
	fmt.Fprintf(w, "<h1>Información del Sistema Operativo</h1>")
	fmt.Fprintf(w, "<strong>Sistema Operativo:</strong> %s<br>", runtime.GOOS)
	fmt.Fprintf(w, "<strong>Arquitectura:</strong> %s<br>", runtime.GOARCH)
	fmt.Fprintf(w, "<strong>Número de CPUs:</strong> %d<br>", runtime.NumCPU())

	d, _ := disk.Usage("/")
	fmt.Fprintf(w, "<strong>Espacio total del disco:</strong> %v GB<br>", d.Total/1024/1024/1024)
	fmt.Fprintf(w, "<strong>Espacio usado del disco:</strong> %v GB (%.2f%%)<br>", d.Used/1024/1024/1024, d.UsedPercent)

	v, _ := mem.VirtualMemory()
	fmt.Fprintf(w, "<strong>Total memoria RAM:</strong> %v GB<br>", v.Total/1024/1024/1024)
	fmt.Fprintf(w, "<strong>Memoria RAM usada:</strong> %v GB (%.2f%%)<br>", v.Used/1024/1024/1024, v.UsedPercent)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(w, "<strong>Error obteniendo el directorio del usuario:</strong> %s<br>", err)
	} else {
		fmt.Fprintf(w, "<strong>Carpeta del Usuario:</strong> %s<br>", homeDir)
	}
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Servidor iniciado en http://localhost:%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
