package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/streadway/amqp"
)

func ReplaceTurkishChars(str string) string {
	replacer := strings.NewReplacer(
		"ç", "c",
		"ğ", "g",
		"ı", "i",
		"ö", "o",
		"ş", "s",
		"ü", "u",
		"Ç", "C",
		"Ğ", "G",
		"İ", "I",
		"Ö", "O",
		"Ş", "S",
		"Ü", "U",
	)
	return replacer.Replace(str)
}

func GetPdf(order Order, db *sql.DB) {
	m := GetMaroto(order, db)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Klasörün var olup olmadığını kontrol edin, yoksa oluşturun
	outputDir := "pdfs"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModePerm)
	}

	fileName := fmt.Sprintf("%s.pdf", order.Date)
	filePath := fmt.Sprintf("%s/%s", outputDir, fileName)

	err = document.Save(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	Publisher(order.Date)
}
func GetMaroto(order Order, db *sql.DB) core.Maroto {
	cfg := config.NewBuilder().
		//WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	//longText := "This is a longer sentence that will be broken into multiple lines \n" +
	//  "as it does not fit into the column otherwise."

	m.AddRow(40,
		image.NewFromFileCol(3, "../img/logo.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		col.New(1),
		col.New(8).Add(
			text.New("E-FATURA", props.Text{
				Size:  25,
				Align: align.Center,
				Style: fontstyle.BoldItalic,
			}),
			text.New("Sevk Irsaliye No 451-592", props.Text{
				Right: 13,
				Top:   14,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Center,
			}),
			text.New(fmt.Sprint(time.Now().Day(), "/", int(time.Now().Month()), "/", time.Now().Year()), props.Text{
				Right: 32,
				Top:   18,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Center,
				//Color: &props.Color{},
			}),
		),
	)
	m.AddRow(5,
		col.New(3).Add(
			text.New("YTÜ Teknopark C1 Blok", props.Text{
				Size:  7,
				Align: align.Center,
				Style: fontstyle.BoldItalic,
			}),
		),
	)

	m.AddRow(1, line.NewCol(12))

	m.AddRow(50,
		col.New(12).Add(
			text.New("Sayin", props.Text{
				Left: 10,
				Size: 15,
				//Align: align.Center,
				Style: fontstyle.BoldItalic,
				Align: align.Left,
				Top:   3,
			}),
			text.New(fmt.Sprint(ReplaceTurkishChars(order.Address.Name), " ", ReplaceTurkishChars(order.Address.Surname)), props.Text{
				Left:  10,
				Top:   10,
				Size:  8,
				Align: align.Left,
			}),
			text.New(fmt.Sprint(ReplaceTurkishChars(order.Address.Adres), " ",
				ReplaceTurkishChars(order.Address.Ilce), " ", ReplaceTurkishChars(order.Address.City)), props.Text{
				Left:  10,
				Top:   28,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				//Color: &props.Color{},
			}),
			text.New("Vergi Dairesi", props.Text{
				Left:  10,
				Top:   35,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				//Color: &props.Color{},
			}),
			text.New("Güngören", props.Text{
				Left:  10,
				Top:   40,
				Size:  8,
				Align: align.Left,
				//Color: &props.Color{},
			}),
		),
	)

	m.AddRow(7,
		text.NewCol(12, "Ürünler", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: &props.Color{185, 185, 185}})

	m.AddRow(20,
		text.NewCol(6, "isim", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.Color{0, 0, 0},
		}),
		text.NewCol(3, "Miktar", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.Color{0, 0, 0},
		}),
		text.NewCol(3, "Fiyat", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.Color{0, 0, 0},
		}),
	)
	var totalPrice int
	for _, Basket := range order.Baskets {

		basket_id, err := strconv.Atoi(Basket.Id)

		if err != nil {
			panic(err.Error())
		}

		product := DBSearch(basket_id, db)
		amount, err := strconv.Atoi(Basket.Amount)
		if err != nil {
			panic(err.Error())
		}
		amount = amount * int(product.Price)
		totalPrice += amount
		m.AddRow(7,
			text.NewCol(3, ReplaceTurkishChars(product.Model), props.Text{
				Top:   1.5,
				Size:  9,
				Style: fontstyle.Bold,
				Align: align.Center,
				Color: &props.Color{0, 0, 0},
			}),
			text.NewCol(3, Basket.Amount, props.Text{
				Top:   1.5,
				Size:  9,
				Style: fontstyle.Bold,
				Align: align.Center,
				Color: &props.Color{0, 0, 0},
			}),
			text.NewCol(3, strconv.Itoa(amount), props.Text{
				Top:   1.5,
				Size:  9,
				Style: fontstyle.Bold,
				Align: align.Center,
				Color: &props.Color{0, 0, 0},
			}),
		)
	}

	m.AddRow(7,
		text.NewCol(12, fmt.Sprint("Toplam Ucret = ", totalPrice), props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.Color{0, 0, 0},
		}))
	return m
}

func Publisher(Date string) {
	conn, err := amqp.Dial("destination rabbitmq")
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

	email := EmailMessage{
		To:      "sending mail",
		Subject: "E-Trade Fatura",
		Body: "Yaptığınız alışverişin faturası aşağıdaki pdf'te bulunmaktadır." +
			"Bizi tercih ettiğiniz için teşekkür ederiz.",
		Date: Date,
	}

	body, err := json.Marshal(email)
	failOnError(err, "Failed to marshal email message")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	fmt.Printf(" [x] Sent %s\n", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
