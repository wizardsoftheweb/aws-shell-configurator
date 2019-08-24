package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func tidyPath(pathComponents ...string) (string, error) {
	return filepath.Abs(filepath.Join(pathComponents...))
}

func EnsureDirectoryExists(pathComponents ...string) error {
	combinedPath, err := tidyPath(pathComponents...)
	if nil == err {
		err = os.MkdirAll(combinedPath, os.ModePerm)
	}
	return err
}

func LoadFile(pathComponents ...string) (string, error) {
	combinedPath, err := tidyPath(pathComponents...)
	if nil != err {
		return "", err
	}
	rawContents, err := ioutil.ReadFile(combinedPath)
	return string(rawContents), err
}
