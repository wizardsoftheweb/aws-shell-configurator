package main

import (
	"regexp"
	"strings"
)

var (
	PatternSectionTitle = regexp.MustCompile(
		`^\[(?:profile\s+)?['"]?([^'"\]]+)['"]?\s*]`,
	)
	PatternSectionKeyValue = regexp.MustCompile(
		`^([^\[:=]+?)\s*[:=]\s*(.+?)\s*$`,
	)
	DotFileConfig = NewDotFile(
		"~/.aws/config",
		"AWS_CONFIG_FILE",
	)
	DotFileCredentials = NewDotFile(
		"~/.aws/credentials",
		"AWS_SHARED_CREDENTIALS_FILE",
	)
)

type AwsDotFile struct {
	DefaultPath         string
	EnvironmentVariable string
	CurrentPath         string
	RawContents         []byte
	Contents            []string
	Profiles            map[string]*AwsProfile
}

func NewDotFile(defaultPath, envVariable string) *AwsDotFile {
	dotFile := &AwsDotFile{
		DefaultPath:         defaultPath,
		EnvironmentVariable: envVariable,
		Profiles:            make(map[string]*AwsProfile),
	}
	dotFile.discoverLocation()
	return dotFile
}

func (f *AwsDotFile) discoverLocation() {
	f.CurrentPath, _ = tidyPath(GetEnvWithDefault(f.EnvironmentVariable, f.DefaultPath))
}

func (f *AwsDotFile) loadContents() error {
	var err error
	f.RawContents, err = LoadFile(f.CurrentPath)
	return err
}

func (f *AwsDotFile) tidyContents() {
	contents := string(f.RawContents)
	contents = strings.TrimSpace(contents)
	f.Contents = strings.Split(contents, "\n")
}

func (f *AwsDotFile) checkForTitle(line string) string {
	matches := PatternSectionTitle.FindAllStringSubmatch(line, -1)
	if 1 != len(matches) {
		return ""
	}
	return matches[0][1]
}

func (f *AwsDotFile) checkForKeyValue(line string) (string, string) {
	matches := PatternSectionKeyValue.FindAllStringSubmatch(line, -1)
	if 1 != len(matches) {
		return "", ""
	}
	return matches[0][1], matches[0][2]
}

func (f *AwsDotFile) parseTitle(line string, profileName *string) {
	parsedName := f.checkForTitle(line)
	if "" != parsedName {
		*profileName = parsedName
		newProfile := NewAwsProfile()
		newProfile.Profile = *profileName
		f.Profiles[*profileName] = newProfile
	}
}

func (f *AwsDotFile) parseKeyValue(line string, profileName string) {
	if nil != f.Profiles[profileName] {
		key, value := f.checkForKeyValue(line)
		if "" != key {
			_, ok := f.Profiles[profileName].Settings[key]
			if ok {
				f.Profiles[profileName].Settings[key].Set(value)
			}
		}
	}
}

func (f *AwsDotFile) parse() {
	var profileName string
	for _, line := range f.Contents {
		f.parseTitle(line, &profileName)
		f.parseKeyValue(line, profileName)
	}
}

func (f *AwsDotFile) LoadAndParse() (map[string]*AwsProfile, error) {
	err := f.loadContents()
	if nil != err {
		return nil, err
	}
	f.tidyContents()
	f.parse()
	return f.Profiles, nil
}

func MergeConfigAndCredentialsProfiles(configProfiles, credentialsProfiles map[string]*AwsProfile) map[string]*AwsProfile {
	for credProfileName, credProfile := range credentialsProfiles {
		credentialsSettings := credProfile.ExtractCredentialsSettings()
		_, ok := configProfiles[credProfileName]
		if !ok {
			configProfiles[credProfileName] = NewAwsProfile()
			configProfiles[credProfileName].Profile = credProfileName
		}
		configProfiles[credProfileName].UpdateSettingValues(credentialsSettings)
		configProfiles[credProfileName].updateFromEnvironment()
	}
	return configProfiles
}
