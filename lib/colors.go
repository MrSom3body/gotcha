package lib

import "fmt"

var Colors = struct {
	Reset   string
	Black   string
	Red     string
	Green   string
	Blue    string
	Yellow  string
	Magenta string
	Cyan    string
	White   string
}{
	Reset:   "\033[0m",
	Black:   "\033[30m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Blue:    "\033[34m",
	Yellow:  "\033[33m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",
}

func GetColors() string {
	return fmt.Sprintf("%s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s",
		Colors.Black,
		Colors.Reset,
		Colors.Red,
		Colors.Reset,
		Colors.Green,
		Colors.Reset,
		Colors.Yellow,
		Colors.Reset,
		Colors.Blue,
		Colors.Reset,
		Colors.Magenta,
		Colors.Reset,
		Colors.Cyan,
		Colors.Reset,
		Colors.White,
		Colors.Reset,
	)
}
