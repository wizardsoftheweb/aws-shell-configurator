package main

import (
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
