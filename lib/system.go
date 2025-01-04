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
		if strings.HasPrefix(scanner.Text(), "PRETTY_NAME=") {
			return strings.Trim(scanner.Text()[13:], "\"")
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
	lastSlash := strings.LastIndex(shell, "/")
	if lastSlash == -1 {
		return shell
	}
	return shell[lastSlash+1:]
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
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			memoryTotal, err = strconv.ParseFloat(strings.Fields(line)[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.HasPrefix(line, "MemAvailable:") {
			memoryAvailable, err = strconv.ParseFloat(strings.Fields(line)[1], 64)
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

func GetIpAddress(ifaceName string) string {
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		log.Fatal(err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ip := ipNet.IP.To4(); ip != nil {
				return ip.String()
			}
		}
	}

	return "no ip on `" + ifaceName + "`"
}
