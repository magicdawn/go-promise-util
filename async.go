package putil

import (
	"github.com/chebyrash/promise"
)

// Async: wrap a func & use fn's return value to resolve, like `co()`
func Async(fn func() interface{}) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		result := fn()
		resolve(result)
	})
}

// AsyncFactory: create a async function, like `co.wrap`
func AsyncFactory(fn func(args ...interface{}) interface{}) func(args ...interface{}) *promise.Promise {
	return func(args ...interface{}) *promise.Promise {
		return Async(func() interface{} {
			return fn(args...)
		})
	}
}
