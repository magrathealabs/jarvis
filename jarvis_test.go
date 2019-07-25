package main

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type JarvisSuite struct {
	gspec.Suite
}

func (suite *JarvisSuite) TestMain() {
	suite.Panics(main)
}

func TestJarvisSuite(t *testing.T) {
	gspec.Run(t, new(JarvisSuite))
}
