# text-generator
Fast text generator on a mask.

Written in Golang. I do not use regular expressions and the fastest. I covered tests and simple! Supporting recursive text generation rules.

[![Build Status](https://travis-ci.org/liderman/text-generator.svg?branch=master)](https://travis-ci.org/liderman/text-generator)&nbsp;[![GoDoc](https://godoc.org/github.com/liderman/text-generator?status.svg)](https://godoc.org/github.com/liderman/text-generator)

Installation
-----------
Stable release (v1):

    go get gopkg.in/liderman/text-generator.v1

Non stable release (master branch):

	go get github.com/liderman/text-generator

Usage
-----------
An example of a simple template text generation:
```go
tg := text_generator.New()
template := "Good {morning|day}!"

fmt.Print(tg.Generate(template))
// Displays: Good morning!

fmt.Print(tg.Generate(template))
// Displays: Good day!
```

An example of a complex generation template text:
```go
tg := text_generator.New()
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

* Need at least `go1.2` or newer.

Documentation
-----------

You can read package documentation [here](http:godoc.org/github.com/liderman/text-generator).

Testing
-----------
Unit-tests:
```bash
go test -v
```

Benchmarks:
```bash
go test -test.bench .
```
The test result on computer mac-mini 2012 (Intel Core i5):
```
PASS
BenchmarkGenerateEasyText-4      1000000              1699 ns/op
BenchmarkGenerateComplexText-4    200000              7430 ns/op
ok      github.com/liderman/text-generator            3.391s
```
