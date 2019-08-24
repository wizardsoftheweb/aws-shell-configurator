package main

import "fmt"

var SettingsInCredentialsFile = []string{
	"aws_access_key_id",
	"aws_secret_access_key",
	"aws_session_token",
}

type AwsProfile struct {
	Profile  string                 `json:"profile"`
	Settings map[string]*AwsSetting `json:"settings"`
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

func (p *AwsProfile) WriteProfile() error {
	credentials := make(map[string]string)
	for _, key := range SettingsInCredentialsFile {
		setting, ok := p.Settings[key]
		if ok && "" != setting.Value {
			credentials[key] = setting.Value
		}
		delete(p.Settings, key)
	}
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
