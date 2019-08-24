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
}

func nilErrorOrPanic(err error) {
	panic(err)
}
