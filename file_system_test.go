package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	. "gopkg.in/check.v1"
)

type FileSystemSuite struct {
	BaseSuite
	fileContents []byte
	fileToWrite  string
}

var _ = Suite(&FileSystemSuite{
	BaseSuite{sharedErrorMessage: "shared file error"},
	[]byte("test data"),
	"",
})

func (s *FileSystemSuite) BrokenPathTidier(input ...string) (string, error) {
	return "", errors.New(s.sharedErrorMessage)
}

func (s *FileSystemSuite) SetUpTest(c *C) {
	s.fileToWrite = filepath.Join(s.workingDirectory, "test.file")
}

func (s *FileSystemSuite) TearDownTest(c *C) {
	_ = os.Remove(s.fileToWrite)
	pathTidier = tidyPath
}

func (s *FileSystemSuite) TestTidyPath(c *C) {
	var tidyPathData = [][]string{
		{"/", "/"},
		{"/some/dir", "/some", "dir"},
		{fmt.Sprintf("%s/%s", s.workingDirectory, "some/dir"), "some", "dir"},
	}
	for _, value := range tidyPathData {
		result, err := tidyPath(value[1:]...)
		c.Assert(err, Not(ErrorMatches), "*")
		c.Assert(result, Equals, value[0])
	}
}

func (s *FileSystemSuite) TestEnsureDirectoryExistsWorksWithCwd(c *C) {
	err := EnsureDirectoryExists(s.workingDirectory)
	c.Assert(err, IsNil)
}

func (s *FileSystemSuite) TestEnsureDirectoryExistsCreatesDirectories(c *C) {
	additionalPathComponents := []string{"some", "dir"}
	fullPath := filepath.Join(
		append(
			[]string{s.workingDirectory},
			additionalPathComponents...,
		)...,
	)
	err := EnsureDirectoryExists(additionalPathComponents...)
	c.Assert(err, IsNil)
	_, err = os.Stat(fullPath)
	c.Assert(os.IsNotExist(err), Equals, false)
}

func (s *FileSystemSuite) TestLoadFilePathError(c *C) {
	pathTidier = s.BrokenPathTidier
	_, err := LoadFile(s.workingDirectory)
	c.Assert(err, ErrorMatches, s.sharedErrorMessage)
}

func (s *FileSystemSuite) TestLoadFileThatDoesntExist(c *C) {
	_, err := LoadFile(filepath.Join(s.workingDirectory, "random", "file"))
	c.Assert(os.IsNotExist(err), Equals, true)
}

func (s *FileSystemSuite) TestLoadFileNonEmpty(c *C) {
	_, filename, _, _ := runtime.Caller(0)
	contents, err := LoadFile(filename)
	c.Assert(err, IsNil)
	c.Assert(contents, Not(Equals), []byte{})
}

func (s *FileSystemSuite) TestWriteFilePathError(c *C) {
	pathTidier = s.BrokenPathTidier
	err := writeFile(s.fileContents, 0600, s.workingDirectory)
	c.Assert(err, ErrorMatches, s.sharedErrorMessage)
}

func (s *FileSystemSuite) TestWriteFileSuccess(c *C) {
	perm := os.FileMode(uint32(0640))
	err := writeFile(s.fileContents, perm, s.fileToWrite)
	c.Assert(err, IsNil)
	contents, err := ioutil.ReadFile(s.fileToWrite)
	c.Assert(err, IsNil)
	c.Assert(contents, DeepEquals, s.fileContents)
	stats, _ := os.Stat(s.fileToWrite)
	c.Assert(stats.Mode().String(), Equals, perm.String())
}

func (s *FileSystemSuite) TestWriteDotFileSuccess(c *C) {
	err := WriteDotFile(s.fileContents, s.fileToWrite)
	c.Assert(err, IsNil)
}
