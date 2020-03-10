package promiseUtil

import (
	"github.com/chebyrash/promise"
)

func Async(fn func() interface{}) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		result := fn()
		resolve(result)
	})
}

func AsyncFactory(fn func(args ...interface{}) interface{}) func(args ...interface{}) *promise.Promise {
	return func(args ...interface{}) *promise.Promise {
		return Async(func() interface{} {
			return fn(args...)
		})
	}
}
