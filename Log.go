package goshell

import (
	"fmt"
)

func Log(args ...interface{}) {
	fmt.Printf("%+v\n", args)
}
