package main

import (
	"regexp"
)

var (
	DotFileSectionTitleLinePattern    = regexp.MustCompile(`^\s*\[(?P<title>[^]]+)].*$`)
	DotFileSectionKeyValueLinePattern = regexp.MustCompilePOSIX(`^(?P<indent>\s*)(?P<key>[^:=[]+?)\s*?(?P<assignment>[:=])\s*(?P<value>.*?)(\s+[;#]\s*(?P<comment>.*?)\s*)?$`)
)

type DotFileSection struct {
	Title  string
	Values *map[string]string
}

func NewDotFileSection(title string, values *map[string]string) *DotFileSection {
	return &DotFileSection{
		Title:  title,
		Values: values,
	}
}
