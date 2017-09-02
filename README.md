# base91-go
[![Travis](https://img.shields.io/travis/Equim-chan/base91-go.svg)](https://travis-ci.org/Equim-chan/base91-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Equim-chan/base91-go)](https://goreportcard.com/report/github.com/Equim-chan/base91-go)
[![license](https://img.shields.io/badge/BSD-3-blue.svg)](https://github.com/Equim-chan/base91-go/blob/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/ekyu.moe/base91)

basE91 codec implemented in Golang.

Migrated from the C and PHP version of Joachim Henke's [basE91](http://base91.sourceforge.net/).

[Nodejs version](https://github.com/Equim-chan/base91)

## Install
```bash
$ go get -u ekyu.moe/base91
```

## Example
```go
package main

import (
    "fmt"

    "ekyu.moe/base91"
)

func main() {
    fmt.Println(base91.EncodeToString([]byte("Hello, 世界"))) //=> >OwJh>}AFU~PUh%Y
    fmt.Println(string(base91.DecodeString(">OwJh>}AFU~PUh%Y"))) //=> Hello, 世界
}
```

Check [godoc](https://godoc.org/ekyu.moe/base91) for further documents.

A CLI utility is also available with `go get ekyu.moe/base91/cmd/base91`.

## Benchmark
Note: 1 op = 1 MB input
```plain
$ go test -bench . -benchmem ekyu.moe/base91
goos: windows
goarch: amd64
pkg: ekyu.moe/base91
BenchmarkDecode-4             20          77053375 ns/op         7157760 B/op         36 allocs/op
BenchmarkDecoder-4            10         146889760 ns/op         6332510 B/op        916 allocs/op
BenchmarkEncode-4            200           6279833 ns/op         8673280 B/op         37 allocs/op
BenchmarkEncoder-4           200           6278146 ns/op         2324512 B/op         24 allocs/op
PASS
ok      ekyu.moe/base91 7.598s
```

(Approximately 159.28 MB/s for encoder and 6.81 MB/s for decoder)

## License
[BSD-3-clause](https://github.com/Equim-chan/base91-go/blob/master/LICENSE)
