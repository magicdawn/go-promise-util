package main

import (
	"fmt"
	"time"

	"github.com/chebyrash/promise"
	promiseUtil "github.com/magicdawn/go-promise-util"
)

func main() {
	sleepUseNewPromise := func(sec int) *promise.Promise {
		return promise.New(func(resolve func(interface{}), reject func(error)) {
			time.Sleep(time.Second * time.Duration(sec))
			resolve(sec * sec)
		})
	}

	sleepUseAsync := func(sec int) *promise.Promise {
		return promiseUtil.Async(func() interface{} {
			time.Sleep(time.Second * time.Duration(sec))
			return sec
		})
	}

	sleepUseAsyncFactory := promiseUtil.AsyncFactory(func(args ...interface{}) interface{} {
		sec := args[0].(int)
		time.Sleep(time.Second * time.Duration(sec))
		return sec
	})

	arr := []interface{}{1, 2, 3, 4, 5}

	p1 := promiseUtil.Map(arr, func(item interface{}, index int, items []interface{}) *promise.Promise {
		return sleepUseNewPromise(item.(int))
	}, 2)
	p2 := promiseUtil.Map(arr, func(item interface{}, index int, items []interface{}) *promise.Promise {
		return sleepUseAsync(item.(int))
	}, 2)
	p3 := promiseUtil.Map(arr, func(item interface{}, index int, items []interface{}) *promise.Promise {
		return sleepUseAsyncFactory(item.(int))
	}, 2)
	
	result, err := promise.All(p1, p2, p3).Await()
	fmt.Println(result, err)
}
