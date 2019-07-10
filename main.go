package main

import (
	"flag"

	gomail "gopkg.in/gomail.v2"
)

func main() {

	mailsFlag := flag.String("to", "", "mail recipient, if multiple recipients use comma as separator ','")
	attachmentFlag := flag.String("attach", "", "files you want to attach, if multiple files use comma as separator ','")
	subjectFlag := flag.String("subj", "Automated mail sended from mailCLI app", "subject of mail message")
	configFlag := flag.String("config", "config/config.json", "location of json config file")
	flag.Parse()

	msg := gomail.NewMessage()
	msg.SetHeader("From", "sendmailrmichalsiwik@gmail.com")

	addRecipients(*mailsFlag, msg)
	addAttachments(*attachmentFlag, msg)

	//msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	msg.SetHeader("Subject", *subjectFlag)
	msg.SetBody("text/html", "Cześć,\nTen mail został wysłany automatycznie za pomocą aplikacji mailCLI w razie jakichkolkwiek pytań, lub problemów proszę o kontakt z <b>michalsiwik@gmail.com</b>!")

	config := readConfig(*configFlag)

	d := gomail.NewDialer(config.Mailserver.SMTPGate, config.Mailserver.SMTPPort, config.Authentication.Login, config.Authentication.Password)

	if err := d.DialAndSend(msg); err != nil {
		panic(err)
	}

}