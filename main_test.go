package main

import (
	. "gopkg.in/check.v1"
)

type RootMainSuite struct {
	BaseSuite
}

var _ = Suite(&RootMainSuite{})

func (s *RootMainSuite) TestRootMain(c *C) {
	c.Assert(
		func() {
			main()
		},
		Not(Panics),
		"",
	)
}
