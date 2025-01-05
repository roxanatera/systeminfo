package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const hostRoot = "/" // Ruta base local

// SystemInfo representa la información del sistema
type SystemInfo struct {
	HostInfo struct {
		OS       string `json:"os"`
		Platform string `json:"platform"`
		Version  string `json:"version"`
		Hostname string `json:"hostname"`
		Uptime   uint64 `json:"uptime"`
	} `json:"host_info"`
	CPUInfo struct {
		ModelName string  `json:"model_name"`
		Cores     int32   `json:"cores"`
		Usage     float64 `json:"usage"`
	} `json:"cpu_info"`
	MemoryInfo struct {
		TotalGB float64 `json:"total_gb"`
		UsedGB  float64 `json:"used_gb"`
		Usage   float64 `json:"usage_percent"`
	} `json:"memory_info"`
	DiskInfo struct {
		TotalGB float64 `json:"total_gb"`
		UsedGB  float64 `json:"used_gb"`
		Usage   float64 `json:"usage_percent"`
	} `json:"disk_info"`
}

// getSystemInfo recopila la información del sistema local
func getSystemInfo() (SystemInfo, error) {
	var info SystemInfo

	// Información del host
	hostInfo, err := host.InfoWithContext(context.TODO())
	if err != nil {
		return info, err
	}
	info.HostInfo.OS = hostInfo.OS
	info.HostInfo.Platform = hostInfo.Platform
	info.HostInfo.Version = hostInfo.PlatformVersion
	info.HostInfo.Hostname = hostInfo.Hostname
	info.HostInfo.Uptime = hostInfo.Uptime

	// Información de CPU
	cpuInfo, err := cpu.InfoWithContext(context.TODO())
	if err != nil {
		return info, err
	}
	if len(cpuInfo) > 0 {
		info.CPUInfo.ModelName = cpuInfo[0].ModelName
		info.CPUInfo.Cores = cpuInfo[0].Cores
	}
	cpuUsage, err := cpu.PercentWithContext(context.TODO(), 0, false)
	if err != nil {
		return info, err
	}
	if len(cpuUsage) > 0 {
		info.CPUInfo.Usage = cpuUsage[0]
	}

	// Información de memoria
	memInfo, err := mem.VirtualMemoryWithContext(context.TODO())
	if err != nil {
		return info, err
	}
	info.MemoryInfo.TotalGB = float64(memInfo.Total) / (1024 * 1024 * 1024)
	info.MemoryInfo.UsedGB = float64(memInfo.Used) / (1024 * 1024 * 1024)
	info.MemoryInfo.Usage = memInfo.UsedPercent

	// Información del disco
	diskInfo, err := disk.Usage(hostRoot)
	if err != nil {
		return info, err
	}
	info.DiskInfo.TotalGB = float64(diskInfo.Total) / (1024 * 1024 * 1024)
	info.DiskInfo.UsedGB = float64(diskInfo.Used) / (1024 * 1024 * 1024)
	info.DiskInfo.Usage = diskInfo.UsedPercent

	return info, nil
}

// saveToJSON guarda la información en un archivo JSON
func saveToJSON(info SystemInfo, filename string) error {
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

func main() {
	// Crear la información del sistema
	info, err := getSystemInfo()
	if err != nil {
		log.Fatalf("Error obteniendo información del sistema: %v\n", err)
	}

	// Guardar la información en un archivo JSON
	filename := "system_info.json"
	err = saveToJSON(info, filename)
	if err != nil {
		log.Fatalf("Error escribiendo JSON en archivo: %v\n", err)
	}
	log.Printf("Información del sistema guardada en '%s'\n", filename)

	// Iniciar el servidor HTTP
	app := fiber.New()

	// Servir archivos estáticos
	app.Static("/css", "./css")
	app.Static("/js", "./js")

	// Servir el archivo index.html desde la raíz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})

	// Ruta para descargar el archivo JSON
	app.Get("/download-json", func(c *fiber.Ctx) error {
		return c.Download(filename)
	})

	// Detectar puerto dinámicamente
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Iniciar el servidor
	log.Printf("Servidor corriendo en http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
