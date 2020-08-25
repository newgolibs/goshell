package goshell

import (
	"fmt"
)

func Logok(args ...interface{}) interface{} {
	fmt.Printf("\033[01;32m%+v\033[0m\n", args)
	return nil
}
