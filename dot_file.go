package main

type DotFile struct {
	Path     string
	Sections []*DotFileSection
}

func NewDotFile(path string, sections []*DotFileSection) *DotFile {
	return &DotFile{
		Path:     path,
		Sections: sections,
	}
}
