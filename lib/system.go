package lib

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

func GetDistribution() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PRETTY_NAME") {
			return strings.Trim(strings.Split(scanner.Text(), "=")[1], "\"")
		}
	}

	return "Unknown"
}

func GetKernel() string {
	content, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(content))
}

func plural(value int) string {
	if value > 1 {
		return "s"
	} else {
		return ""
	}
}

func GetUptime() string {
	rawContent, err := os.ReadFile("/proc/uptime")
	if err != nil {
		log.Fatal(err)
	}

	uptime, err := strconv.ParseFloat(strings.Fields(string(rawContent))[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	parts := make([]string, 0, 3)
	total_minutes := int(math.Round(uptime / 60))
	days := total_minutes / (60 * 24)
	hours := (total_minutes % (60 * 24)) / 60
	minutes := total_minutes % 60

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d day%s", days, plural(days)))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hour%s", hours, plural(hours)))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d minute%s", minutes, plural(minutes)))
	}

	return strings.Join(parts, " ")
}

func GetShell() string {
	shell := os.Getenv("SHELL")
	shellPath := strings.Split(shell, "/")
	shell = shellPath[len(shellPath)-1]
	return shell
}

func GetMemory() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	memoryTotal := 0.0
	memoryAvailable := 0.0
	memoryUsed := 0.0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		switch line[0] {
		case "MemTotal:":
			memoryTotal, err = strconv.ParseFloat(line[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case "MemAvailable:":
			memoryAvailable, err = strconv.ParseFloat(line[1], 64)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	memoryUsed = memoryTotal - memoryAvailable
	memoryTotal = memoryTotal / 1024 / 1024
	memoryUsed = memoryUsed / 1024 / 1024

	return fmt.Sprintf("%.1f GiB / %.1f GiB", memoryUsed, memoryTotal)
}

func GetIpAddress() string {
	int, err := net.InterfaceByName("wlp2s0")
	if err != nil {
		log.Fatal(err)
	}
	item, err := int.Addrs()
	if err != nil {
		log.Fatal(err)
	}

	var ip net.IP
	for _, addr := range item {
		switch v := addr.(type) {
		case *net.IPNet:
			if !v.IP.IsLoopback() {
				if v.IP.To4() != nil {
					ip = v.IP
				}
			}
		}
	}
	if ip != nil {
		return ip.String()
	} else {
		return ""
	}
}
