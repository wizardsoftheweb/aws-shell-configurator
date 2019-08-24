package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func tidyPath(pathComponents ...string) (string, error) {
	return filepath.Abs(filepath.Join(pathComponents...))
}

var (
	pathTidier    = tidyPath
	dotFileWriter = func(contents []byte, pathComponents ...string) error {
		return writeFile(contents, 0600, pathComponents...)
	}
)

func EnsureDirectoryExists(pathComponents ...string) error {
	combinedPath, err := pathTidier(pathComponents...)
	if nil == err {
		err = os.MkdirAll(combinedPath, os.ModePerm)
	}
	return err
}

func LoadFile(pathComponents ...string) (string, error) {
	combinedPath, err := pathTidier(pathComponents...)
	if nil != err {
		return "", err
	}
	rawContents, err := ioutil.ReadFile(combinedPath)
	return string(rawContents), err
}

func writeFile(contents []byte, permissions os.FileMode, pathComponents ...string) error {
	combinedPath, err := pathTidier(pathComponents...)
	if nil != err {
		return err
	}
	return ioutil.WriteFile(combinedPath, contents, permissions)
}

func WriteDotFile(contents []byte, pathComponents ...string) error {
	return dotFileWriter(contents, pathComponents...)
}
