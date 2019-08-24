package main

import (
	. "gopkg.in/check.v1"
)

type AwsSettingSuite struct {
	BaseSuite
	awsAccessKeyId AwsSetting
	output         AwsSetting
}

var _ = Suite(&AwsSettingSuite{})

func (s *AwsSettingSuite) SetUpTest(c *C) {
	s.awsAccessKeyId = AwsSetting{
		EnvironmentVariable: "AWS_ACCESS_KEY_ID",
	}
	s.output = AwsSetting{
		EnvironmentVariable: "AWS_DEFAULT_OUTPUT",
		Value:               "json",
		AllowedValues:       []string{"json", "table", "text"},
	}
}

func (s *AwsSettingSuite) TestSetAllowedSetting(c *C) {
	c.Assert(s.awsAccessKeyId.Value, Equals, "")
	s.awsAccessKeyId.Set(s.output.Value)
	c.Assert(s.awsAccessKeyId.Value, Equals, s.output.Value)
	c.Assert(s.output.Value, Equals, "json")
	s.output.Set("table")
	c.Assert(s.output.Value, Equals, "table")
}

func (s *AwsSettingSuite) TestSetUnallowed(c *C) {
	c.Assert(s.output.Value, Equals, "json")
	s.output.Set("qqq")
	c.Assert(s.output.Value, Equals, "json")
}
