package main

import (
	"fmt"
	"os"
	"strings"

	. "gopkg.in/check.v1"
)

type AwsDotFileSuite struct {
	BaseSuite
	dotFile      AwsDotFile
	fileWasSet   bool
	existingFile string
	defaultPath  string
	contents     string
	profileName  string
	settingName  string
}

var _ = Suite(&AwsDotFileSuite{})

func (s *AwsDotFileSuite) SetUpTest(c *C) {
	s.existingFile, s.fileWasSet = os.LookupEnv(DotFileCredentials.EnvironmentVariable)
	_ = os.Unsetenv(DotFileCredentials.EnvironmentVariable)
	s.defaultPath, _ = tidyPath(DotFileCredentials.DefaultPath)
	s.dotFile = AwsDotFile{
		DefaultPath:         DotFileCredentials.DefaultPath,
		EnvironmentVariable: DotFileCredentials.EnvironmentVariable,
		Profiles:            make(map[string]*AwsProfile),
	}
	s.contents = "this is a test string"
	s.profileName = "default"
	s.settingName = "test"
}

func (s *AwsDotFileSuite) TearDownTest(c *C) {
	if s.fileWasSet {
		_ = os.Setenv(DotFileCredentials.EnvironmentVariable, s.existingFile)
	} else {
		_ = os.Unsetenv(DotFileCredentials.EnvironmentVariable)
	}
}

func (s *AwsDotFileSuite) TestNewDotFile(c *C) {
	newDotFile := NewDotFile(DotFileCredentials.DefaultPath, DotFileCredentials.EnvironmentVariable)
	c.Assert(newDotFile.CurrentPath, Equals, s.defaultPath)
}

func (s *AwsDotFileSuite) TestDiscoverLocation(c *C) {
	c.Assert(s.dotFile.CurrentPath, Equals, "")
	s.dotFile.discoverLocation()
	c.Assert(s.dotFile.CurrentPath, Equals, s.defaultPath)
}

func (s *AwsDotFileSuite) TestLoadContents(c *C) {
	c.Assert(s.dotFile.RawContents, IsNil)
	s.dotFile.CurrentPath = s.currentFilename
	err := s.dotFile.loadContents()
	c.Assert(err, IsNil)
	c.Assert(s.dotFile.RawContents, Not(IsNil))
}

func (s *AwsDotFileSuite) TestTidyContents(c *C) {
	s.dotFile.RawContents = []byte(s.contents)
	c.Assert(s.dotFile.Contents, HasLen, 0)
	s.dotFile.tidyContents()
	c.Assert(s.dotFile.Contents, HasLen, 1)
	c.Assert(s.dotFile.Contents[0], Equals, s.contents)
}

func (s *AwsDotFileSuite) TestCheckForTitleSuccess(c *C) {
	matrix := []struct {
		profile string
		result  string
	}{
		{"default", "default"},
		{`profile "test"`, "test"},
		{"this won't work", ""},
	}
	for _, entry := range matrix {
		line := fmt.Sprintf("[%s]", entry.profile)
		c.Assert(s.dotFile.checkForTitle(line), Equals, entry.result)
	}
}

func (s *AwsDotFileSuite) TestCheckForKeyValue(c *C) {
	matrix := []struct {
		input string
		key   string
		value string
	}{
		{
			"key = value", "key", "value",
		},
		{
			"key = value\nkey2=value2", "", "",
		},
		{
			"# comment", "", "",
		},
		{
			"[default]", "", "",
		},
	}
	for _, entry := range matrix {
		key, value := s.dotFile.checkForKeyValue(entry.input)
		c.Assert(key, Equals, entry.key)
		c.Assert(value, Equals, entry.value)
	}
}

func (s *AwsDotFileSuite) TestParseTitle(c *C) {
	matrix := []struct {
		line    string
		profile string
		length  int
	}{
		{"[default]", "default", 1},
	}
	for _, entry := range matrix {
		var profileName string
		c.Assert(s.dotFile.Profiles, HasLen, 0)
		s.dotFile.parseTitle(entry.line, &profileName)
		c.Assert(s.dotFile.Profiles, HasLen, entry.length)
		c.Assert(entry.profile, Equals, profileName)
	}
}

func (s *AwsDotFileSuite) TestParseKeyValue(c *C) {
	s.dotFile.Profiles = map[string]*AwsProfile{
		s.profileName: {
			Profile: s.profileName,
			Settings: map[string]*AwsSetting{
				s.settingName: {
					Value: s.defaultPath,
				},
			},
		},
	}
	c.Assert(
		s.dotFile.Profiles[s.profileName].Settings[s.settingName].Value,
		Equals,
		s.defaultPath,
	)
	s.dotFile.parseKeyValue(
		fmt.Sprintf("%s = %s", s.settingName, s.contents),
		s.profileName,
	)
	c.Assert(
		s.dotFile.Profiles[s.profileName].Settings[s.settingName].Value,
		Equals,
		s.contents,
	)
	s.dotFile.parseKeyValue(
		"[default]",
		s.profileName,
	)
	c.Assert(
		s.dotFile.Profiles[s.profileName].Settings[s.settingName].Value,
		Equals,
		s.contents,
	)
}

func (s *AwsDotFileSuite) TestParse(c *C) {
	s.dotFile.Contents = strings.Split(
		fmt.Sprintf(
			"[%s]\n%s = %s",
			s.profileName,
			s.settingName,
			s.defaultPath,
		),
		"\n",
	)
	c.Assert(s.dotFile.Profiles, HasLen, 0)
	s.dotFile.parse()
	c.Assert(s.dotFile.Profiles, HasLen, 1)
}

func (s *AwsDotFileSuite) TestLoadAndParseSuccess(c *C) {
	s.dotFile.CurrentPath = s.currentFilename
	_, err := s.dotFile.LoadAndParse()
	c.Assert(err, IsNil)
}

func (s *AwsDotFileSuite) TestLoadAndParseFileDne(c *C) {
	s.dotFile.CurrentPath = s.currentFilename[1:]
	_, err := s.dotFile.LoadAndParse()
	c.Assert(err, ErrorMatches, "*")
}

func (s *AwsDotFileSuite) TestMergeConfigAndCredentialsProfiles(c *C) {
	config := map[string]*AwsProfile{
		"default": {
			Profile: "default",
			Settings: map[string]*AwsSetting{
				"aws_access_key_id": {
					Value: "config",
				},
				"output": {
					Value: "config",
				},
			},
		},
	}
	creds := map[string]*AwsProfile{
		"default": {
			Profile: "default",
			Settings: map[string]*AwsSetting{
				"aws_access_key_id": {
					Value: "cred",
				},
				"output": {
					Value: "cred",
				},
			},
		},
		"test": {
			Profile: "test",
			Settings: map[string]*AwsSetting{
				"aws_access_key_id": {
					Value: "cred",
				},
				"output": {
					Value: "cred",
				},
			},
		},
	}
	profiles := MergeConfigAndCredentialsProfiles(config, creds)
	c.Assert(
		profiles["default"].Settings["aws_access_key_id"].Value,
		Equals,
		"cred",
	)
	c.Assert(
		profiles["default"].Settings["output"].Value,
		Equals,
		"config",
	)
	c.Assert(
		profiles["test"].Settings["aws_access_key_id"].Value,
		Equals,
		"cred",
	)
	c.Assert(
		profiles["test"].Settings["output"].Value,
		Equals,
		"json",
	)
}
