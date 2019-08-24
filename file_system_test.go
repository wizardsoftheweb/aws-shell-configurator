package main

import (
	"fmt"

	. "gopkg.in/check.v1"
)

type FileSystemSuite struct {
	BaseSuite
}

var _ = Suite(&FileSystemSuite{
	BaseSuite{SharedErrorMessage: "shared file error"},
})

func (s *FileSystemSuite) TestTidyPath(c *C) {
	var tidyPathData = [][]string{
		{"/", "/"},
		{"/some/dir", "/some", "dir"},
		{fmt.Sprintf("%s/%s", s.WorkingDirectory, "some/dir"), "some", "dir"},
	}
	for _, value := range tidyPathData {
		result, err := tidyPath(value[1:]...)
		c.Assert(err, Not(ErrorMatches), "*")
		c.Assert(result, Equals, value[0])
	}
}
