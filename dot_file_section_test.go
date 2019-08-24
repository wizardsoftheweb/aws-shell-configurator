package main

import (
	. "gopkg.in/check.v1"
)

type DotFileSectionSuite struct {
	BaseSuite
	title  string
	values *map[string]string
}

var _ = Suite(&DotFileSectionSuite{})

func (s *DotFileSectionSuite) SetUpTest(c *C) {
	s.title = "default"
	values := make(map[string]string)
	s.values = &values
}

func (s *DotFileSectionSuite) TestNewDotFileSection(c *C) {
	dotFileSection := NewDotFileSection(s.title, s.values)
	c.Assert(dotFileSection.Title, Equals, s.title)
	c.Assert(dotFileSection.Values, DeepEquals, s.values)
}
