package main

import (
	"fmt"
	mailer "github.com/audetv/smtp-testing"
	"log"
	"time"
)

func main() {
	config := mailer.MailerConfig{
		Host:     "localhost",
		Port:     1025,
		Username: "",
		Password: "ss",
		Timeout:  5 * time.Second,
		Sender:   "johndoe@example.com",
	}

	sender := mailer.New(config)
	err := sender.Send("johndoe@example.com", "welcome.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email successful sent!")
}
