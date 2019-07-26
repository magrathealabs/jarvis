package main

import (
	"os"
	"testing"

	"github.com/krakenlab/gspec"
)

type JarvisSuite struct {
	gspec.Suite
}

func (suite *JarvisSuite) TestMain() {
	os.Setenv("WEB_PORT", "-1")
	os.Setenv("WEB_INTERFACE", "localhost")

	suite.Panics(main)
}

func TestJarvisSuite(t *testing.T) {
	gspec.Run(t, new(JarvisSuite))
}
