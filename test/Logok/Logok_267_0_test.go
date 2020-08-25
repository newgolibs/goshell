package Logok

import (
	"github.com/newgolibs/goshell"
	"github.com/newgolibs/goshell/test/Logok/Logok_267_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Logok_267_0 struct {
}

func TestLogok_Logok_267_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range Logok_267_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Logok_267_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Logok_267_0) run(input, arg interface{}) interface{} {
	goshell.Logok(input)
	return nil
}
