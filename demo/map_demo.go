package main

import (
	"fmt"
	"time"

	"github.com/chebyrash/promise"
	promiseUtil "github.com/magicdawn/go-promise-util"
)

func handle(item interface{}, index int, items []interface{}) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		time.Sleep(time.Second * time.Duration(item.(int)))
		resolve(item.(int) * item.(int))
	})
}

func main() {
	arr := []interface{}{1, 2, 3, 4, 5}
	p := promiseUtil.Map(arr, handle, 1)
	result, err := p.Await()
	fmt.Println(result, err)
}
