package putil

import (
	"testing"
	"time"

	"github.com/chebyrash/promise"
)

func handle(item interface{}, index int, items []interface{}) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		val := item.(int)
		time.Sleep(time.Millisecond * 10 * time.Duration(val))
		resolve(val * val)
	})
}

func TestMap(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	p := Map(arr, handle, 2)
	res, err := p.Await()

	if err != nil {
		t.Errorf("err is not empty")
	}

	result := res.([]interface{})

	for i := 0; i < len(arr); i++ {
		val := (arr[i]).(int)
		if result[i] != val*val {
			t.Errorf("unexpected result")
		}
	}

}
