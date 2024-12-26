package lib

import "fmt"

var Colors = struct {
	Reset   string
	Red     string
	Green   string
	Blue    string
	Yellow  string
	Magenta string
	Cyan    string
	White   string
}{
	Reset:   "\033[0m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Blue:    "\033[34m",
	Yellow:  "\033[33m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",
}

func GetColors() string {
	return fmt.Sprintf("%s  %s  %s  %s  %s  %s  %s%s",
		Colors.Red,
		Colors.Green,
		Colors.Yellow,
		Colors.Blue,
		Colors.Magenta,
		Colors.Cyan,
		Colors.White,
		Colors.Reset,
	)
}
