package main

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BaseSuite struct {
	sharedErrorMessage string
	workingDirectory   string
}

var _ = Suite(&BaseSuite{})

var currentworkingDirectory, _ = os.Getwd()

func (s *BaseSuite) SetUpSuite(c *C) {
	s.workingDirectory = c.MkDir()
	_ = os.Chdir(s.workingDirectory)
}

func (s *BaseSuite) TearDownSuite(c *C) {
	_ = os.Chdir(currentworkingDirectory)
}
