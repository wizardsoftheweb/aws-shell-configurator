package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BaseSuite struct {
	SharedErrorMessage string
}

var _ = Suite(&BaseSuite{})
