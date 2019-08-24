package main

import (
	"os"

	. "gopkg.in/check.v1"
)

type AwsProfileSuite struct {
	BaseSuite
	profile         AwsProfile
	profileWasSet   bool
	existingProfile string
}

var _ = Suite(&AwsProfileSuite{})

func (s *AwsProfileSuite) SetUpTest(c *C) {
	s.profile = AwsProfile{}
	s.existingProfile, s.profileWasSet = os.LookupEnv("AWS_PROFILE")
	_ = os.Unsetenv("AWS_PROFILE")
}

func (s *AwsProfileSuite) TearDownTest(c *C) {
	if s.profileWasSet {
		_ = os.Setenv("AWS_PROFILE", s.existingProfile)
	}
}

func (s *AwsProfileSuite) TestNewAwsProfile(c *C) {
	profile := NewAwsProfile()
	c.Assert(profile.Settings["output"].Value, Equals, "json")
}
