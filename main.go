package main

import (
    "fmt"
    "os"
	"net/http"
    "runtime"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, `<html><head><title>Página Inicial</title></head><body>
        <h1>Bienvenido a la Página de Información del Sistema</h1>
        <p><a href="/systeminfo">Ver la información del sistema del servidor aquí</a></p>
        <h2>Información del Sistema del Cliente</h2>
        <div id="infoCliente"></div>
        <script>
        document.getElementById('infoCliente').innerHTML = '<strong>Sistema Operativo:</strong> ' + navigator.platform + 
        '<br><strong>Navegador:</strong> ' + navigator.userAgent;
        </script>
        </body></html>`)
}


func systemInfo(w http.ResponseWriter, r *http.Request) {
    // Establece el tipo de contenido y la codificación de caracteres
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    // Comienza a generar la página HTML con la información del sistema
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
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Maneja la ruta /systeminfo
	http.HandleFunc("/systeminfo", systemInfo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}
	fmt.Printf("Servidor iniciado en http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

