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
	os.Setenv("SERVER_PORT", "-1")
	suite.Panics(main)
}

func TestJarvisSuite(t *testing.T) {
	gspec.Run(t, new(JarvisSuite))
}
