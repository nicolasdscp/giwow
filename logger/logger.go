package logger

import (
	"fmt"
	"os"
)

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

var Fatal = func(format string, a ...any) {
	fmt.Printf("> "+format+"\n", a...)
	os.Exit(1)
}

var Print = func(format string, a ...any) {
	fmt.Printf("> "+format+"\n", a...)
}

var Verbose = func(format string, a ...any) {
	fmt.Printf("> [v] "+format+"\n", a...)
}

var Debug = func(format string, a ...any) {
	fmt.Printf("> [d] "+format+"\n", a...)
}

var empty = func(_ string, _ ...any) {}
