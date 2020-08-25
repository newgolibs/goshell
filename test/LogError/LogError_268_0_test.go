package LogError

import (
	"github.com/newgolibs/goshell"
	"github.com/newgolibs/goshell/test/LogError/LogError_268_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type LogError_268_0 struct {
}

func TestLogError_LogError_268_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range LogError_268_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], LogError_268_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (LogError_268_0) run(input, arg interface{}) interface{} {
	goshell.LogError(input)
	return nil
}
