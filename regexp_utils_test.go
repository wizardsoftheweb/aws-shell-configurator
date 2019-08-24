package main

import (
	"regexp"
	"strings"

	. "gopkg.in/check.v1"
)

type RegexpUtilsSuite struct {
	BaseSuite
	pattern *regexp.Regexp
}

var _ = Suite(&RegexpUtilsSuite{})

func (s *RegexpUtilsSuite) SetUpTest(c *C) {
	s.pattern = regexp.MustCompile(`^(?P<key>[^=]+)=(?P<value>.*)$`)
}

func (s *RegexpUtilsSuite) TestRegexpSubmatchNamed(c *C) {
	searchStringSlice := []string{"key", "value"}
	results := RegexpSubmatchNamed(s.pattern, strings.Join(searchStringSlice, "="))
	c.Assert(results["key"], Equals, searchStringSlice[0])
	c.Assert(results["value"], Equals, searchStringSlice[1])
}

func (s *RegexpUtilsSuite) TestGetSpecificKeysSuccess(c *C) {
	input := make(map[string]string)
	input["one"] = "two"
	keys := []string{"one"}
	results, err := getSpecificMapKeys(keys, input)
	c.Assert(err, IsNil)
	c.Assert(results["one"], Equals, input["one"])
}

func (s *RegexpUtilsSuite) TestGetSpecificKeysMissingKey(c *C) {
	input := make(map[string]string)
	input["one"] = "two"
	keys := []string{"qqq"}
	_, err := getSpecificMapKeys(keys, input)
	c.Assert(err, ErrorMatches, "missing the following keys.*")
}
