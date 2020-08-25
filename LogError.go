package goshell

import (
	"fmt"
)

func LogError(args ...interface{}) {
	fmt.Printf("\033[01;31m%+v\033[0m\n", args)
}
