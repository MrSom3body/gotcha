package lib

import "fmt"

var Colors = struct {
	Reset  string
	Values []string
}{
	Reset: "\033[0m",
	Values: []string{
		"\033[30m", // Black
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
		"\033[37m", // White
	},
}

func GetColors() string {
	result := ""
	for _, color := range Colors.Values {
		result += fmt.Sprintf("%s %s ", color, Colors.Reset)
	}
	return result
}
