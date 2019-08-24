package main

import "fmt"

func main() {
	fmt.Println("rad")
	// configProfiles, _ := DotFileConfig.LoadAndParse()
	// credsProfiles, _ := DotFileCredentials.LoadAndParse()
	// allProfiles := MergeConfigAndCredentialsProfiles(configProfiles, credsProfiles)
	// for _, profile := range allProfiles {
	// 	_ = profile.compileProfile()
	// }
	newProfile := NewAwsProfile()
	for key, setting := range newProfile.Settings {
		fmt.Printf("%s: %#v\n", key, setting)
	}
	_ = newProfile.compileProfile()
}

func nilErrorOrPanic(err error) {
	panic(err)
}
