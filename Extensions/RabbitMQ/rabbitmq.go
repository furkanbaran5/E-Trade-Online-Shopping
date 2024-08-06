package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
)

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Date    string `json:"date"`
}

func main() {
	Consume()
}

func Consume() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_queue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var email EmailMessage
			err := json.Unmarshal(d.Body, &email)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}
			fmt.Printf("Received a message: %s\n", d.Body)
			err = sendEmail(email)
			if err != nil {
				log.Printf("Failed to send email: %s", err)
			} else {
				fmt.Println("Email sent successfully")
			}
		}
	}()

	fmt.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendEmail(email EmailMessage) error {
	pdfPath := "../proje/proje/src/go/pdfs/" + email.Date + ".pdf"
	from := "etradecompany00@gmail.com"
	password := "demvohruznlhwioe"

	// SMTP sunucu ayarları
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Mesaj oluşturma
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)

	// PDF dosyasını ekle
	m.Attach(pdfPath)

	// E-posta gönderme
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	// PDF dosyasını sil
	if err := os.Remove(pdfPath); err != nil {
		return err
	}

	return nil
}
