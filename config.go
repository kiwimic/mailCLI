package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config struct which definies configuration needed for sending mails
type Config struct {
	Mailserver struct {
		SMTPGate string `json:"smtp_gate"`
		SMTPPort int    `json:"smtp_port"`
		FromMail string `json:"from_mail"`
	} `json:"mailserver"`
	Authentication struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	} `json:"authentication"`
}

func readConfig(configJSON string) Config {

	var config Config

	//Reading JSON
	// Open our jsonFile
	jsonFile, err := os.Open(configJSON)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Unmarshaling
	json.Unmarshal(byteValue, &config)

	return config
}


