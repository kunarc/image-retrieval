package runtime

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func PrintPanic(baseError *error) {
	if r := recover(); r != nil {
		err := r.(error)
		*baseError = errors.New(err.Error())
		fmt.Printf("panic: %v\n", r)
		debug.PrintStack()
	}
}
