package goshell

import (
	"fmt"
)

func Log(args ...interface{}) interface{} {
	fmt.Printf("%+v\n", args)
	return nil
}
