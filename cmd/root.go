package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MrSom3body/gotcha/lib"
	"github.com/spf13/cobra"
)

var (
	ifaceName string = "lo"
	version   string
	color     string = "6"
)

var rootCmd = &cobra.Command{
	Use:     "gotcha",
	Version: version,
	Short:   "gotcha - a fast and small fetch tool for Linux 🐹",
	Long: `gotcha is a speedy and minimalistic fetch tool written in Go. 
It provides essential system information such as distribution, kernel version, 
uptime, shell, desktop environment/window manager, memory usage, and local IP.
Customization is minimal by design, focusing on simplicity and speed.`,
	Run: func(cmd *cobra.Command, args []string) {
		colorInt, _ := strconv.Atoi(color)
		keyColor := lib.Colors.Values[colorInt]
		format := `
 %s󰌽  Distro    %s  %s
 %s  Kernel    %s  %s
 %s  Uptime    %s  %s
 %s  Shell     %s  %s
 %s󰧨  DE/WM     %s  %s
 %s  Memory    %s  %s
 %s  Local IP  %s  %s

 %s
`

		fmt.Printf(format,
			keyColor, lib.Colors.Reset, lib.GetDistribution(),
			keyColor, lib.Colors.Reset, lib.GetKernel(),
			keyColor, lib.Colors.Reset, lib.GetUptime(),
			keyColor, lib.Colors.Reset, lib.GetShell(),
			keyColor, lib.Colors.Reset, lib.GetDesktopEnvironment(),
			keyColor, lib.Colors.Reset, lib.GetMemory(),
			keyColor, lib.Colors.Reset, lib.GetIpAddress(),
			lib.GetColors())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
