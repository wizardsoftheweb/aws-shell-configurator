package main

import (
	"errors"

	. "gopkg.in/check.v1"
)

type RootMainSuite struct {
	BaseSuite
}

var _ = Suite(&RootMainSuite{
	BaseSuite{
		sharedErrorMessage: "main error",
	},
})

func (s *RootMainSuite) TestRootMain(c *C) {
	c.Assert(
		func() {
			main()
		},
		Not(Panics),
		"*",
	)
}

func (s *RootMainSuite) TestNilErrorOrPanic(c *C) {
	c.Assert(
		func() {
			nilErrorOrPanic(errors.New(s.sharedErrorMessage))
		},
		PanicMatches,
		s.sharedErrorMessage,
	)
}
