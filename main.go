package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

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

// getSystemInfo genera información del sistema para cada solicitud
func getSystemInfo() (SystemInfo, error) {
	var info SystemInfo

	// Información del host
	hostInfo, err := host.Info()
	if err != nil {
		return info, err
	}
	info.HostInfo.OS = hostInfo.OS
	info.HostInfo.Platform = hostInfo.Platform
	info.HostInfo.Version = hostInfo.PlatformVersion
	info.HostInfo.Hostname = hostInfo.Hostname
	info.HostInfo.Uptime = hostInfo.Uptime

	// Información de CPU
	cpuInfo, err := cpu.Info()
	if err != nil {
		return info, err
	}
	if len(cpuInfo) > 0 {
		info.CPUInfo.ModelName = cpuInfo[0].ModelName
		info.CPUInfo.Cores = cpuInfo[0].Cores
	}

	// Información de memoria
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return info, err
	}
	info.MemoryInfo.TotalGB = float64(memInfo.Total) / (1024 * 1024 * 1024)
	info.MemoryInfo.UsedGB = float64(memInfo.Used) / (1024 * 1024 * 1024)
	info.MemoryInfo.Usage = memInfo.UsedPercent

	// Información del disco
	diskInfo, err := disk.Usage("/")
	if err != nil {
		return info, err
	}
	info.DiskInfo.TotalGB = float64(diskInfo.Total) / (1024 * 1024 * 1024)
	info.DiskInfo.UsedGB = float64(diskInfo.Used) / (1024 * 1024 * 1024)
	info.DiskInfo.Usage = diskInfo.UsedPercent

	return info, nil
}

func main() {
	// Iniciar el servidor HTTP
	app := fiber.New()

	// Ruta para generar y descargar el archivo JSON
	app.Get("/generate-json", func(c *fiber.Ctx) error {
		// Generar la información del sistema
		info, err := getSystemInfo()
		if err != nil {
			return c.Status(500).SendString("Error obteniendo información del sistema")
		}

		// Crear el archivo JSON temporalmente
		outputPath := filepath.Join(os.TempDir(), "system_info.json")
		file, err := os.Create(outputPath)
		if err != nil {
			return c.Status(500).SendString("Error creando archivo JSON")
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(info); err != nil {
			return c.Status(500).SendString("Error escribiendo información en archivo JSON")
		}

		// Descargar el archivo JSON
		return c.Download(outputPath, "system_info.json")
	})

	// Servir página estática
	app.Static("/", "./dist")

	// Puerto dinámico para producción
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Servidor corriendo en http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
