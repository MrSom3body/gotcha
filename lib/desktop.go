package lib

import (
	"fmt"
	"os"
)

func GetDesktopEnvironment() string {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	session := os.Getenv("XDG_SESSION_TYPE")
	return fmt.Sprintf("%s (%s)", desktop, session)
}
