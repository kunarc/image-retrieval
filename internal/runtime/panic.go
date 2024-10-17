package runtime

import (
	"errors"
	"runtime/debug"
)

func PrintPanic(baseError *error) {
	if r := recover(); r != nil {
		if val, ok := r.(error); ok {
			*baseError = errors.New(val.Error())
		} else {
			*baseError = errors.New(r.(string))
		}
		debug.PrintStack()
	}
}
