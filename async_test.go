package putil

import (
	"testing"
	"time"

	"github.com/chebyrash/promise"
)

func TestAsync(t *testing.T) {
	sleepUseNewPromise := func(sec int) *promise.Promise {
		return promise.New(func(resolve func(interface{}), reject func(error)) {
			time.Sleep(time.Second * time.Duration(sec))
			resolve(sec * sec)
		})
	}

	sleepUseAsync := func(sec int) *promise.Promise {
		return Async(func() interface{} {
			time.Sleep(time.Second * time.Duration(sec))
			return sec + 1
		})
	}

	sleepUseAsyncFactory := AsyncFactory(func(args ...interface{}) interface{} {
		sec := args[0].(int)
		time.Sleep(time.Second * time.Duration(sec))
		return sec + 2
	})

	p1 := sleepUseNewPromise(1)
	p2 := sleepUseAsync(1)
	p3 := sleepUseAsyncFactory(1)

	res, err := promise.All(p1, p2, p3).Await()

	if err != nil {
		t.Errorf("err is not empty")
	}

	result := res.([]interface{})
	if result[0] != 1 || result[1] != 2 || result[2] != 3 {
		t.Errorf("unexpected result")
	}
}
