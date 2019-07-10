package main

import (
	"flag"
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {

	mailsFlag := flag.String("to", "", "mail recipient, if multiple recipients use comma as separator ','")
	attachmentFlag := flag.String("attach", "", "files you want to attach, if multiple files use comma as separator ','")
	subjectFlag := flag.String("subj", "Automated mail sended from mailCLI app", "subject of mail message")
	configFlag := flag.String("config", "C:/goworkspace/src/github.com/mailCLI/config/config.json", "location of json config file")
	checkConfFlag := flag.Bool("check", false, "This command checks if path to config and log have existing files")
	flag.Parse()

	if *checkConfFlag {
		fmt.Println(checkConfiguration(*configFlag, "C:/goworkspace/src/github.com/mailCLI/logs/log.txt"))
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "sendmailrmichalsiwik@gmail.com")

	addRecipients(*mailsFlag, msg)
	addAttachments(*attachmentFlag, msg)

	//msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	msg.SetHeader("Subject", *subjectFlag)
	msg.SetBody("text/html", "Cześć,\nTen mail został wysłany automatycznie za pomocą aplikacji mailCLI w razie jakichkolkwiek pytań, lub problemów proszę o kontakt z <b>michalsiwik@gmail.com</b>!")

	config := readConfig(*configFlag)

	d := gomail.NewDialer(config.Mailserver.SMTPGate, config.Mailserver.SMTPPort, config.Authentication.Login, config.Authentication.Password)

	logOperation(config.Mailserver.FromMail, *mailsFlag, *attachmentFlag, *subjectFlag, config.Mailserver.SMTPGate, config.Mailserver.SMTPPort, *configFlag, "C:/goworkspace/src/github.com/mailCLI/logs/log.txt")

	if err := d.DialAndSend(msg); err != nil {
		panic(err)
	}

}
