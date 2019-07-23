# gspec

[![Build Status](https://travis-ci.org/krakenlab/gspec.svg?branch=master)](https://travis-ci.org/krakenlab/gspec)
[![Coverage Status](https://coveralls.io/repos/github/krakenlab/gspec/badge.svg)](https://coveralls.io/github/krakenlab/gspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/krakenlab/gspec)](https://goreportcard.com/report/github.com/krakenlab/gspec)
[![GoDoc](https://godoc.org/github.com/krakenlab/gspec?status.svg)](https://godoc.org/github.com/krakenlab/gspec)

Golang spec library based on testify.

```golang
package hello

import "github.com/krakenlab/gspec"

type MySuite struct {
    gspec.Suite
}

func (suite *MySuite) SetupTest() {
    // Do anything
    suite.Suite.SetupTest()
}

func (suite *MySuite) TestTwoPlusTwo() {
    suite.Equal(2+2, 4)
}

func TestMySuite(t *testing.T) {
    Run(t, new(MySuite))
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
    git clone git@github.com:krakenlab/gspec.git
```
