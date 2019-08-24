package main

import (
	"regexp"
	"strings"

	. "gopkg.in/check.v1"
)

type RegexpUtilsSuite struct {
	BaseSuite
}

var _ = Suite(&RegexpUtilsSuite{})

func (s *RegexpUtilsSuite) TestRegexpSubmatchNamed(c *C) {
	pattern := regexp.MustCompile(`^(?P<key>[^=]+)=(?P<value>.*)$`)
	searchStringSlice := []string{"key", "value"}
	results := RegexpSubmatchNamed(pattern, strings.Join(searchStringSlice, "="))
	c.Assert(results["key"], Equals, searchStringSlice[0])
	c.Assert(results["value"], Equals, searchStringSlice[1])
}
