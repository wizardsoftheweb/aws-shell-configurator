package main

import (
	"os"
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
	s.WorkingDirectory = c.MkDir()
	_ = os.Chdir(s.WorkingDirectory)
}
