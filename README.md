# base91-go
[![Travis](https://img.shields.io/travis/Equim-chan/base91-go.svg)](https://travis-ci.org/Equim-chan/base91-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Equim-chan/base91-go)](https://goreportcard.com/report/github.com/Equim-chan/base91-go)
[![license](https://img.shields.io/badge/BSD-3.0-blue.svg)](https://github.com/Equim-chan/base91-go/blob/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/ekyu.moe/base91)

basE91 codec implemented in Golang.

Migrated from the C and PHP version of Joachim Henke's [basE91](http://base91.sourceforge.net/).

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

## License
[BSD-3.0](https://github.com/Equim-chan/base91-go/blob/master/LICENSE)
