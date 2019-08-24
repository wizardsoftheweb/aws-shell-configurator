package main

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
