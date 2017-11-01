package main

import (
	"flag"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

func main() {
	var from, to, cc, subject, body, file, smtp, user, pass string
	var smtpPort int
	flag.StringVar(&from, "from", "", "set from")
	flag.StringVar(&to, "to", "", "set to")
	flag.StringVar(&cc, "cc", "", "set cc")
	flag.StringVar(&subject, "subject", "", "set subject")
	flag.StringVar(&body, "body", "", "set body")
	flag.StringVar(&file, "file", "", "set file attach")
	flag.StringVar(&user, "user", "", "set username")
	flag.StringVar(&pass, "pass", "", "set password")
	flag.StringVar(&smtp, "smtp", "", "set smtp server")
	flag.IntVar(&smtpPort, "smtpport", 25, "set smtp port")
	flag.Parse()

	m := gomail.NewMessage()
	m.SetHeader("From", from)

	tos := strings.Split(to, ",")
	m.SetHeader("To", tos...)

	m.SetAddressHeader("Cc", cc, cc)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	if file != "" {
		m.Attach(file)
	}

	d := gomail.NewDialer(smtp, smtpPort, user, pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}
