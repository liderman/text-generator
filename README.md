# text-generator
Fast text generator on a mask.

[![Go Report Card](https://goreportcard.com/badge/github.com/liderman/text-generator)](https://goreportcard.com/report/github.com/liderman/text-generator)
[![codecov](https://codecov.io/gh/liderman/text-generator/branch/master/graph/badge.svg?token=53NH949TQY)](https://codecov.io/gh/liderman/text-generator)
[![Go Reference](https://pkg.go.dev/badge/github.com/liderman/text-generator.svg)](https://pkg.go.dev/github.com/liderman/text-generator)
[![Release](https://img.shields.io/github/release/liderman/text-generator.svg?style=flat-square)](https://github.com/liderman/text-generator/releases/latest)


Written in Golang. I do not use regular expressions and the fastest. I covered tests and simple! Supporting recursive text generation rules.

Installation
-----------
	go get github.com/liderman/text-generator

Usage
-----------
An example of a simple template text generation:
```go
tg := textgenerator.New()
template := "Good {morning|day}!"

fmt.Print(tg.Generate(template))
// Displays: Good morning!

fmt.Print(tg.Generate(template))
// Displays: Good day!
```

An example of a complex generation template text:
```go
tg := textgenerator.New()
template := "{Good {morning|evening|day}|Goodnight|Hello}, {friend|brother}! {How are you|What's new with you}?"

fmt.Print(tg.Generate(template))
// Displays: Good morning, friend! How are you?

fmt.Print(tg.Generate(template))
// Displays: Good day, brother! What's new with you?

fmt.Print(tg.Generate(template))
// Displays: Hello, friend! How are you?
...
```

Requirements
-----------

* Need at least `go1.14` or newer.

Features
-----------

* Without external dependencies
* Fast (does not use regular expressions)
* Simple
* Supports caching
* Covered with tests

Documentation
-----------

You can read package documentation [here](https://pkg.go.dev/github.com/liderman/text-generator).

Testing
-----------
Unit-tests:
```bash
go test -v -race ./...
```

Run linter:
```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.40 golangci-lint run -v
```

Benchmarks
-----------
```bash
go test -test.bench . -benchmem -benchtime=10s
```
The test result on computer MacBook Pro 2020 (Intel Core i5):
```
goos: darwin
goarch: amd64
pkg: github.com/liderman/text-generator
BenchmarkGenerateEasyText-8             22446540               531 ns/op             200 B/op          8 allocs/op
BenchmarkGenerateComplexText-8           4721838              2552 ns/op            1351 B/op         24 allocs/op
PASS
ok      github.com/liderman/text-generator      27.227s
```

CONTRIBUTE
-----------
* write code
* run `go fmt ./...`
* run all linters and tests (see above)
* create a PR describing the changes

LICENSE
-----------
MIT

AUTHOR
-----------
Konstantin Osipov <k.osipov.msk@gmail.com>