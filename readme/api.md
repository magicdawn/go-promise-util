## API

### Map

```go
Map(
  items interface{}[],
  fn func(item interface{}, index int, items []interface{}) *promise.Promise,
  concurrency int
) *promise.Promise
```

like node.js `promise.map` / `async.parallelLimit`

### Async & AsyncFactory

```go
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
```