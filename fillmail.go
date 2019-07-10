package main

import (
	"fmt"
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
