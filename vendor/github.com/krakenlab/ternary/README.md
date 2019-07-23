# ternary

[![Build Status](https://travis-ci.org/krakenlab/ternary.svg?branch=master)](https://travis-ci.org/krakenlab/ternary)
[![Coverage Status](https://coveralls.io/repos/github/krakenlab/ternary/badge.svg)](https://coveralls.io/github/krakenlab/ternary)
[![Go Report Card](https://goreportcard.com/badge/github.com/krakenlab/ternary)](https://goreportcard.com/report/github.com/krakenlab/ternary)
[![GoDoc](https://godoc.org/github.com/krakenlab/ternary?status.svg)](https://godoc.org/github.com/krakenlab/ternary)

Golang ternary lib.
Write ternary for:

- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64, uintptr
- byte
- rune
- float32, float64
- complex64, complex128
- interface{}
- error
- func

Legacy syntax:

```golang
package main

import (
    "fmt"
)

func main() {
    a := 3
    b := 4

    c := b
    if a > b {
        c = a
    }

    fmt.Println("Biggest value is", c)
}
```

Ternary syntax:

```golang
package main

import (
    "fmt"

    "github.com/krakenlab/ternary"
)

func main() {
    a := 3
    b := 4

    c := ternary.Int(a > b, a, b)
    fmt.Println("Biggest value is", c)
}
```

## Devlopment

Golang 1.12:

```sh
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

    gvm install go1.12 -B
    gvm use go1.12
```

Govendor:

```sh
    go get -u github.com/kardianos/govendor
```

Golint-CI:

```sh
    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
    golangci-lint run
```

GOPATH:

```sh
    mkdir -p $GOPATH/src/github.com/krakenlab/

    cd $GOPATH/src/github.com/krakenlab/
    git clone git@github.com:krakenlab/ternary.git
```
