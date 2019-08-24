package main

import (
	"path/filepath"
	"runtime"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BaseSuite struct {
	SharedErrorMessage string
	WorkingDirectory   string
}

var _ = Suite(&BaseSuite{})

func (s *BaseSuite) SetUpSuite(c *C) {
	_, filename, _, _ := runtime.Caller(0)
	s.WorkingDirectory = filepath.Dir(filename)
}
