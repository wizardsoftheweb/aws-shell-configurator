package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func tidyPath(pathComponents ...string) (string, error) {
	return filepath.Abs(filepath.Join(pathComponents...))
}

var pathTidier = tidyPath

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
