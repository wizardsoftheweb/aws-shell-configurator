package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var SettingsInCredentialsFile = []string{
	"aws_access_key_id",
	"aws_secret_access_key",
	"aws_session_token",
}

type AwsProfile struct {
	Profile  string                 `json:"profile"`
	Settings map[string]*AwsSetting `json:"settings"`
}

func NewAwsProfile() *AwsProfile {
	var profile AwsProfile
	contents, _ := ioutil.ReadFile("default-aws-settings.json")
	_ = json.Unmarshal(contents, &profile)
	return &profile
}

func (p *AwsProfile) isActiveProfile() bool {
	return os.Getenv("AWS_PROFILE") == p.Profile
}

func (p *AwsProfile) updateFromEnvironment() {
	if p.isActiveProfile() {
		for key, setting := range p.Settings {
			if "" != setting.EnvironmentVariable {
				p.Settings[key].Set(GetEnvWithDefault(setting.EnvironmentVariable, setting.Value))
			}
		}
	}
}

func (p *AwsProfile) compileCredentialsFile(profile string, values map[string]string) string {
	output := fmt.Sprintf("[%s]\n", profile)
	for key, value := range values {
		output += fmt.Sprintf("%s = %s\n", key, value)
	}
	return output
}

func (p *AwsProfile) compileBaseConfigFile(profile string, values map[string]string) string {
	if "default" != profile {
		profile = fmt.Sprintf(`profile "%s"`, profile)
	}
	output := fmt.Sprintf("[%s]\n", profile)
	for key, value := range values {
		output += fmt.Sprintf("%s = %s\n", key, value)
	}
	return output
}

func (p *AwsProfile) ExtractCredentialsSettings() map[string]string {
	credentials := make(map[string]string)
	for _, key := range SettingsInCredentialsFile {
		setting, ok := p.Settings[key]
		if ok && "" != setting.Value {
			credentials[key] = setting.Value
		}
		delete(p.Settings, key)
	}
	return credentials
}

func (p *AwsProfile) compileProfile() error {
	credentials := p.ExtractCredentialsSettings()
	config := make(map[string]string)
	for key, setting := range p.Settings {
		if "" != setting.Value {
			config[key] = setting.Value
		}
	}
	credentialsFile := p.compileCredentialsFile(p.Profile, credentials)
	configFile := p.compileBaseConfigFile(p.Profile, config)
	fmt.Println(credentialsFile)
	fmt.Println(configFile)
	return nil
}

func (p *AwsProfile) UpdateSettings(newSettings map[string]*AwsSetting) {
	for key, value := range newSettings {
		p.Settings[key] = value
	}
}

func (p *AwsProfile) UpdateSettingValues(newValues map[string]string) {
	for key, value := range newValues {
		p.Settings[key].Value = value
	}
}
