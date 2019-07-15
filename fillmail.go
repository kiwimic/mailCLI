package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

func addRecipients(mailString string, msg *gomail.Message) {

	mails := strings.Split(mailString, ",")
	addresses := make([]string, len(mails))

	for i, mail := range mails {
		addresses[i] = msg.FormatAddress(mail, "")
	}

	if mailString == "" {
		fmt.Println("You have not added any mail addresss")
		os.Exit(1)
	} else {
		msg.SetHeader("To", addresses...)
	}

}

func addAttachments(attach string, msg *gomail.Message) {

	attachments := strings.Split(attach, ",")
	//fmt.Println(attachments)
	if attach != "" {

		for _, file := range attachments {
			msg.Attach(file)
		}
	} else {
		fmt.Println("You have not attached any files")
	}

}

func addBody(body string, msg *gomail.Message) {
	//fmt.Println("isBodyFile result:", isBodyFile(body))
	if isBodyFile(body) {
		//logic for reading file
		//Filling it with info
		//body = new string with fill

		bodyFile, err := os.Open(body)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer bodyFile.Close()

		// reading bytes
		bodyContent, err := ioutil.ReadAll(bodyFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		body = string(bodyContent)
	}
	//For debugging
	//fmt.Println("body:\n", body)

	msg.SetBody("text/html", body)

}
