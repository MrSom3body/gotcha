package cmd

import (
	"fmt"
	"os"

	"github.com/MrSom3body/gotcha/lib"
	"github.com/spf13/cobra"
)

var IfaceName string = "lo"

var rootCmd = &cobra.Command{
	Use:   "gotcha",
	Short: "gotcha a quick little fetch tool",
	Long:  "gotcha a quick little fetch tool written in go",
	Run: func(cmd *cobra.Command, args []string) {
		keyColor := lib.Colors.Green
		format := `
 %s󰌽  Distro     %s %s
 %s  Kernel     %s %s
 %s  Uptime     %s %s
 %s  Shell      %s %s
 %s󰧨  DE/WM      %s %s
 %s  Memory     %s %s
 %s  Local IP   %s %s

    %s
`

		fmt.Printf(format,
			keyColor,
			lib.Colors.Reset,
			lib.GetDistribution(),
			keyColor,
			lib.Colors.Reset,
			lib.GetKernel(),
			keyColor,
			lib.Colors.Reset,
			lib.GetUptime(),
			keyColor,
			lib.Colors.Reset,
			lib.GetShell(),
			keyColor,
			lib.Colors.Reset,
			lib.GetDesktopEnvironment(),
			keyColor,
			lib.Colors.Reset,
			lib.GetMemory(),
			keyColor,
			lib.Colors.Reset,
			lib.GetIpAddress(IfaceName),
			lib.GetColors())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
