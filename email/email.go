package email

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"os"
	"strconv"

	"github.com/go-mail/mail"
	"github.com/tecolotedev/stori_back/config"
)

// Fill template for signup email
func SendSignupEmail(name string, id int32, to string) {
	f, err := os.ReadFile("email/signup.html")
	if err != nil {
		fmt.Println("err loading template: ", err)
	}
	tmpl, err := template.New("template").Parse(string(f))
	if err != nil {
		fmt.Println(err)
	}

	var bodyContentBuffer bytes.Buffer

	err = tmpl.Execute(&bodyContentBuffer, struct {
		Name      string
		UrlSignup string
	}{
		Name:      name,
		UrlSignup: config.EnvVars.FRONT_URL + "/verifyAccount?id=" + strconv.Itoa(int(id)),
	})
	if err != nil {
		fmt.Println(err)
	}

	SendEmail(to, bodyContentBuffer.String())

}

type Record struct {
	Date        string
	Transaction float64
	Reason      string
}

type TransferSummary struct {
	Month   string
	Type    string
	Average float64
}

// Fill template for report email
func SendReportEmail(to string, accountId int32, balance float64, records []Record, summaries []TransferSummary) {
	f, err := os.ReadFile("email/report.html")
	if err != nil {
		fmt.Println("err loading template: ", err)
	}
	tmpl, err := template.New("template").Parse(string(f))
	if err != nil {
		fmt.Println(err)
	}

	var bodyContentBuffer bytes.Buffer

	err = tmpl.Execute(&bodyContentBuffer, struct {
		LastRecords  []Record
		Balance      float64
		Id           int32
		AllTransfers []TransferSummary
	}{
		LastRecords:  records,
		Balance:      math.Round(balance*100) / 100,
		Id:           accountId,
		AllTransfers: summaries,
	})
	if err != nil {
		fmt.Println(err)
	}

	SendEmail(to, bodyContentBuffer.String())
}

// Send email with mailtrap
func SendEmail(to, htmlContent string) {

	user := config.EnvVars.EMAIL_USER
	password := config.EnvVars.EMAIL_PASSWORD
	host := config.EnvVars.EMAIL_HOST
	port, err := strconv.Atoi(config.EnvVars.EMAIL_PORT)

	if err != nil {
		port = 587
	}

	from := "hello@tecolotedev.com"

	m := mail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", htmlContent)

	d := mail.NewDialer(host, port, user, password)

	if err := d.DialAndSend(m); err != nil {

		fmt.Println("err seding email: ", err)

	}

}
