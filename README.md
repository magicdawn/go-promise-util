# go-promise-util

> promise util for go

[![Build Status](https://img.shields.io/travis/magicdawn/go-promise-util.svg?style=flat-square)](https://travis-ci.org/magicdawn/go-promise-util)
[![Coverage Status](https://img.shields.io/codecov/c/github/magicdawn/go-promise-util.svg?style=flat-square)](https://codecov.io/gh/magicdawn/go-promise-util)
[![GoDoc](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/magicdawn/go-promise-util)
[![Go Report Card](https://goreportcard.com/badge/github.com/magicdawn/go-promise-util?style=flat-square)](https://goreportcard.com/report/github.com/magicdawn/go-promise-util)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

## Install

```sh
$ go get gopkg.in/magicdawn/go-promise-util
```

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

## Changelog

[CHANGELOG.md](CHANGELOG.md)

## License

the MIT License http://magicdawn.mit-license.org
