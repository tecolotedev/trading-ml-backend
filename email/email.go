package email

import (
	"strconv"
	"sync"

	"github.com/go-mail/mail"
	"github.com/tecolotedev/trading-ml-backend/config"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

var pack = "email"

type EmailTemplate = int64

const (
	SignupTemplate EmailTemplate = iota
	ReportTemplate
)

type Mail struct {
	Template EmailTemplate
	Data     map[string]any
	To       string
}

type MailerStruct struct {
	MailChan     chan Mail
	MailDoneChan chan bool
	WG           *sync.WaitGroup
}

func (m *MailerStruct) ListenForEmails() {
	for {
		select {
		case mail := <-m.MailChan:
			switch mail.Template {
			case SignupTemplate:
				name := mail.Data["name"].(string)
				id := mail.Data["id"].(int32)
				to := mail.To
				m.WG.Add(1)
				go SendSignupEmail(name, id, "Verify Your Account", to, m.WG)
			}
		case <-m.MailDoneChan:
			return
		}
	}
}

var Mailer = MailerStruct{
	MailChan:     make(chan Mail),
	MailDoneChan: make(chan bool),
}

// type Record struct {
// 	Date        string
// 	Transaction float64
// 	Reason      string
// }

// type TransferSummary struct {
// 	Month   string
// 	Type    string
// 	Average float64
// }

// // Fill template for report email
// func SendReportEmail(to string, accountId int32, balance float64, records []Record, summaries []TransferSummary) {
// 	f, err := os.ReadFile("email/report.html")
// 	if err != nil {
// 		fmt.Println("err loading template: ", err)
// 	}
// 	tmpl, err := template.New("template").Parse(string(f))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var bodyContentBuffer bytes.Buffer

// 	err = tmpl.Execute(&bodyContentBuffer, struct {
// 		LastRecords  []Record
// 		Balance      float64
// 		Id           int32
// 		AllTransfers []TransferSummary
// 	}{
// 		LastRecords:  records,
// 		Balance:      math.Round(balance*100) / 100,
// 		Id:           accountId,
// 		AllTransfers: summaries,
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	SendEmail(to, bodyContentBuffer.String())
// }

// Send email with mailtrap
func SendEmail(to, subject, htmlContent string) {

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
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlContent)

	d := mail.NewDialer(host, port, user, password)

	if err := d.DialAndSend(m); err != nil {
		utils.Log.ErrorLog(err, pack)
	}

}
