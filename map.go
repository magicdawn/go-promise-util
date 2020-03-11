package promiseUtil

import (
	"sync"

	"github.com/chebyrash/promise"
	. "github.com/visionmedia/go-debug"
)

var debug = Debug("promiseUtil")

// Map : coutil.Map with concurrency
func Map(
	items []interface{},
	fn func(interface{}, int, []interface{}) *promise.Promise,
	concurrency int) *promise.Promise {

	return promise.New(func(resolve func(interface{}), reject func(error)) {
		if concurrency < 0 {
			concurrency = 1
		}

		// control flow
		total := len(items)
		running := 0
		started := 0
		completed := 0
		ret := make([]interface{}, total)
		chComplete := make(chan int, 1)
		chError := make(chan error, total)
		returned := false
		var mu sync.Mutex

		// oncomplete callback
		var oncomplete func()
		oncomplete = func() {
			mu.Lock()
			defer mu.Unlock()

			if returned {
				return
			}

			if completed >= total {
				chComplete <- 1
				return
			}

			for started < total && running < concurrency {
				if returned {
					break
				}

				go func(item interface{}, index int) {
					// new Task
					debug("starting %d", index)

					var err error
					p := fn(item, index, items)
					ret[index], err = p.Await()
					if err != nil {
						chError <- err
						return
					}

					// notify
					mu.Lock()
					running--
					completed++
					mu.Unlock()

					oncomplete()
				}(items[started], started)

				started++
				running++
			}
		}

		go oncomplete()

		select {
		case <-chComplete:
			returned = true
			resolve(ret)
		case err := <-chError:
			returned = true
			reject(err)
		}
	})
}
