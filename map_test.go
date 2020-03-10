package promiseUtil

import (
	"fmt"
	"testing"
	"time"
	"github.com/chebyrash/promise"
)

func handle(item interface{}, index int, items []interface{}) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		time.Sleep(time.Duration(1000 * item.(int)))
		resolve(item.(int) * item.(int))
	})
}

func Test(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	p := Map(arr, handle, 1)
	result, err := p.Await()
	fmt.Println(result, err)
}
