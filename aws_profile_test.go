package main

import (
	"fmt"
	"os"

	. "gopkg.in/check.v1"
)

type AwsProfileSuite struct {
	BaseSuite
	profile          AwsProfile
	profileWasSet    bool
	existingProfile  string
	envVariable      string
	envValue         string
	settingsKey      string
	credsSettingsKey string
}

var _ = Suite(&AwsProfileSuite{})

func (s *AwsProfileSuite) SetUpTest(c *C) {
	s.profile = AwsProfile{}
	s.envVariable = "FAKE_AWS_ACCESS_KEY_ID"
	s.envValue = "qqq"
	s.settingsKey = "aws_access_key_id"
	_ = os.Setenv(s.envVariable, s.envValue)
	s.profile.Settings = map[string]*AwsSetting{
		s.settingsKey: {
			EnvironmentVariable: s.envVariable,
			Value:               "",
			AllowedValues:       []string{},
		},
	}
	s.existingProfile, s.profileWasSet = os.LookupEnv("AWS_PROFILE")
	_ = os.Unsetenv("AWS_PROFILE")
}

func (s *AwsProfileSuite) TearDownTest(c *C) {
	if s.profileWasSet {
		_ = os.Setenv("AWS_PROFILE", s.existingProfile)
	} else {
		_ = os.Unsetenv("AWS_PROFILE")
	}
}

func (s *AwsProfileSuite) TestNewAwsProfile(c *C) {
	profile := NewAwsProfile()
	c.Assert(profile.Settings["output"].Value, Equals, "json")
}

func (s *AwsProfileSuite) TestIsActiveProfile(c *C) {
	s.profile.Profile = "default"
	c.Assert(s.profile.isActiveProfile(), Not(Equals), true)
}

func (s *AwsProfileSuite) TestUpdateFromEnvironment(c *C) {
	c.Assert(s.profile.Settings[s.settingsKey].Value, Not(Equals), s.envValue)
	s.profile.updateFromEnvironment()
	c.Assert(s.profile.Settings[s.settingsKey].Value, Equals, s.envValue)

}

func (s *AwsProfileSuite) TestCompileCredentialsFileDefault(c *C) {
	profileName := "default"
	s.profile.Settings[s.settingsKey].Set(s.envValue)
	output := s.profile.compileCredentialsFile(profileName, s.profile.ExtractCredentialsSettings())
	c.Assert(output, Equals, fmt.Sprintf("[%s]\n%s = %s\n", profileName, s.settingsKey, s.envValue))
}

func (s *AwsProfileSuite) TestCompileBaseConfigFileDefault(c *C) {
	profileName := "default"
	s.profile.Settings[s.settingsKey].Set(s.envValue)
	output := s.profile.compileBaseConfigFile(profileName, s.profile.ExtractCredentialsSettings())
	c.Assert(output, Equals, fmt.Sprintf("[%s]\n%s = %s\n", profileName, s.settingsKey, s.envValue))
}

func (s *AwsProfileSuite) TestCompileBaseConfigFileNotDefault(c *C) {
	profileName := "not default"
	s.profile.Settings[s.settingsKey].Set(s.envValue)
	output := s.profile.compileBaseConfigFile(profileName, s.profile.ExtractCredentialsSettings())
	c.Assert(output, Equals, fmt.Sprintf("[profile \"%s\"]\n%s = %s\n", profileName, s.settingsKey, s.envValue))
}
