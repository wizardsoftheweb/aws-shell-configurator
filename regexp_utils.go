package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
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

func getSpecificMapKeys(keysToGrab []string, input map[string]string) (map[string]string, error) {
	found := make(map[string]string)
	missing := []string{}
	for _, key := range keysToGrab {
		value, ok := input[key]
		if ok {
			found[key] = value
		} else {
			missing = append(missing, key)
		}
	}
	var err error
	if 0 < len(missing) {
		err = errors.New(fmt.Sprintf("Missing the following keys: %s", strings.Join(missing, ", ")))
	}
	return found, err
}
