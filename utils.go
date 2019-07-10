package main

import (
	"flag"
	"fmt"
	"os"
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

func checkConfiguration(conf string, log string) string {

	ret := ""

	if !fileExists(conf) {
		ret = ret + fmt.Sprintf("config not found at: %s", conf) + "\n"
	}

	if !fileExists(log) {
		ret = ret + fmt.Sprintf("log file not found at: %s", log)
	}

	return ret
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
