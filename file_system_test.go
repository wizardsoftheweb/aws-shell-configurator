package main

import (
	"fmt"
	"os"
	"path/filepath"

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

func (s *FileSystemSuite) TestEnsureDirectoryExistsWorksWithCwd(c *C) {
	err := EnsureDirectoryExists(s.WorkingDirectory)
	c.Assert(err, IsNil)
}

func (s *FileSystemSuite) TestEnsureDirectoryExistsCreatesDirectories(c *C) {
	additionalPathComponents := []string{"some", "dir"}
	fullPath := filepath.Join(
		append(
			[]string{s.WorkingDirectory},
			additionalPathComponents...,
		)...,
	)
	err := EnsureDirectoryExists(additionalPathComponents...)
	c.Assert(err, IsNil)
	_, err = os.Stat(fullPath)
	c.Assert(os.IsNotExist(err), Equals, false)
}
