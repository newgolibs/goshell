package Log

import (
	"github.com/newgolibs/goshell"
	"github.com/newgolibs/goshell/test/Log/Log_269_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Log_269_0 struct {
}

func TestLog_Log_269_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range Log_269_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Log_269_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Log_269_0) run(input, arg interface{}) interface{} {
	goshell.Log(input)
	return nil
}
