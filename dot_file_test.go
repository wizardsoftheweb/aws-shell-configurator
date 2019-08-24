package main

import (
	"path/filepath"

	. "gopkg.in/check.v1"
)

type DotFileSuite struct {
	BaseSuite
	path     string
	sections []*DotFileSection
}

var _ = Suite(&DotFileSuite{})

func (s *DotFileSuite) SetUpTest(c *C) {
	s.path = filepath.Join(s.workingDirectory, "dot.file")
	values := make(map[string]string)
	s.sections = []*DotFileSection{
		{
			Title:  "default",
			Values: &values,
		},
	}
}

func (s *DotFileSuite) TestNewDotFile(c *C) {
	dotFile := NewDotFile(s.path, s.sections)
	c.Assert(dotFile.Path, Equals, s.path)
	c.Assert(dotFile.Sections, DeepEquals, s.sections)
}
