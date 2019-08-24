package main

import (
	"regexp"
)

func RegexpSubmatchNamed(regexpPattern *regexp.Regexp, input string) map[string]string {
	discovered := make(map[string]string)
	match := regexpPattern.FindStringSubmatch(input)
	keys := regexpPattern.SubexpNames()
	for index, value := range match {
		discovered[keys[index]] = value
	}
	return discovered
}
