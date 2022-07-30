package logger

import "fmt"

var (
	VerboseMode bool
	DebugMode   bool
)

func Init() {
	if DebugMode {
		return
	} else if VerboseMode {
		Debug = empty
	} else {
		Verbose = empty
		Debug = empty
	}
}

var Print = func(msg string) {
	fmt.Printf("> %s\n", msg)
}

var Verbose = func(msg string) {
	fmt.Printf("[v] %s\n", msg)
}

var Debug = func(msg string) {
	fmt.Printf("[d] %s\n", msg)
}

var empty = func(_ string) {}
