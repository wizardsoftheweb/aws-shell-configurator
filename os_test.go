package main

import (
	"os"

	. "gopkg.in/check.v1"
)

type OsSuite struct {
	BaseSuite
	EnvironmentVariable string
	EnvDefaultValue     string
	EnvSetValue         string
}

var _ = Suite(&OsSuite{
	EnvironmentVariable: "OS_TEST_ENV_VARIABLE",
	EnvDefaultValue:     "this is the default value",
	EnvSetValue:         "qqq",
})

func (s *OsSuite) SetUpTest(c *C) {
	_ = os.Setenv(s.EnvironmentVariable, s.EnvSetValue)
}

func (s *OsSuite) TearDownTest(c *C) {
	_ = os.Unsetenv(s.EnvironmentVariable)
}

func (s *OsSuite) TestGetEnvWithDefaultValueSet(c *C) {
	grabbedValue := GetEnvWithDefault(s.EnvironmentVariable, s.EnvDefaultValue)
	c.Assert(grabbedValue, Not(Equals), s.EnvDefaultValue)
	c.Assert(grabbedValue, Equals, s.EnvSetValue)
}

func (s *OsSuite) TestGetEnvWithDefaultValueUnset(c *C) {
	_ = os.Unsetenv(s.EnvironmentVariable)
	grabbedValue := GetEnvWithDefault(s.EnvironmentVariable, s.EnvDefaultValue)
	c.Assert(grabbedValue, Not(Equals), s.EnvSetValue)
	c.Assert(grabbedValue, Equals, s.EnvDefaultValue)
}
