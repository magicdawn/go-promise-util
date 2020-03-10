# go-promise-util

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