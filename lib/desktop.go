package lib

import (
	"fmt"
	"os"
	"strings"
)

func GetDesktopEnvironment() string {
	desktop := strings.TrimPrefix(os.Getenv("XDG_CURRENT_DESKTOP"), "start-hyprland:")
	session := os.Getenv("XDG_SESSION_TYPE")
	return fmt.Sprintf("%s (%s)", desktop, session)
}
