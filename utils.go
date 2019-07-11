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

func checkConfiguration(conf string, log string) string {

	ret := ""

	boolConf, errConf := fileExists(conf)
	boolLog, errLog := fileExists(log)

	if !boolConf {
		ret = ret + fmt.Sprintf("config not found at: %s", conf) + "\n"
		fmt.Println("config file err:", errConf)
	}

	if !boolLog {
		ret = ret + fmt.Sprintf("log file not found at: %s", log)
		fmt.Println("log file err:", errLog)
	}

	return ret
}

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

//isBodyFile right now because of -body flag I want to check
//if user add path to file, or only string with text to body
//so I used a regex to test if string form body flag have special characters
//I have to test if in not fails with something like C:\\dir1\\dir2
func isBodyFile(s string) bool {

	ret := false
	regextTest, err := regexp.Compile("[<>\\*\\?!@$#\\(\\)\\[\\]]")

	if err != nil {
		fmt.Println("regex complie err for check is body a file: ", err)
	}

	if regextTest.MatchString(s) {
		//fmt.Println("It's not a file (as regex result)")
		ret = false
	} else {
		ret2, err := fileExists(s)
		fmt.Println("err: ", err)
		ret = ret2
	}

	return ret
}
