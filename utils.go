package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func fileExists(name string) bool {

	_, err := os.Open(name)
	if err != nil {
		return false
	}
	return true
}

//isBodyFile right now because of -body flag I want to check
//if user add path to file, or only string with text to body
//so I used a regex to test if string form body flag have special characters
//I have to test if in not fails with something like C:\\dir1\\dir2
func isBodyFile(s string) bool {
	//ret := false
	regextTest, err := regexp.Compile("[<>\\*\\?!@$#\\(\\)\\[\\]]")
	if err != nil {
		fmt.Println("regex complie err for check is body a file: ", err)
	}

	if regextTest.MatchString(s) {
		return false
	}
	return fileExists(s)
}
