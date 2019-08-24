package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	. "gopkg.in/check.v1"
)

type FileSystemSuite struct {
	BaseSuite
}

var _ = Suite(&FileSystemSuite{
	BaseSuite{SharedErrorMessage: "shared file error"},
})

func (s *FileSystemSuite) TearDownTest(c *C) {
	pathTidier = tidyPath
}

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

func (s *FileSystemSuite) TestLoadFileThatDoesntExist(c *C) {
	_, err := LoadFile(filepath.Join(s.WorkingDirectory, "random", "file"))
	c.Assert(os.IsNotExist(err), Equals, true)
}

func (s *FileSystemSuite) TestLoadFileNonEmpty(c *C) {
	_, filename, _, _ := runtime.Caller(0)
	contents, err := LoadFile(filename)
	c.Assert(err, IsNil)
	c.Assert(contents, Not(Equals), "")
}

func (s *FileSystemSuite) TestLoadFilePathError(c *C) {
	pathTidier = func(input ...string) (string, error) {
		return "", errors.New(s.SharedErrorMessage)
	}
	_, err := LoadFile(s.WorkingDirectory)
	c.Assert(err, ErrorMatches, s.SharedErrorMessage)
}
