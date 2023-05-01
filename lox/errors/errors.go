package errors

import (
	"fmt"
	"os"
)

var (
	hadError        bool
	hadRuntimeError bool
)

func Report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	hadError = true
}

func Error(line int, message string) {
	Report(line, "", message)
}
