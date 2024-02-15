package email

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/tecolotedev/trading-ml-backend/config"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

// Fill template for signup email
func SendSignupEmail(name string, id int32, subject, to string, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.ReadFile("email/templates/signup.html")
	if err != nil {
		utils.Log.ErrorLog(err, pack)
	}
	tmpl, err := template.New("template").Parse(string(f))
	if err != nil {
		utils.Log.ErrorLog(err, pack)
	}

	token, err := utils.CreateToken(id, time.Second*20)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
	}

	var bodyContentBuffer bytes.Buffer
	url := fmt.Sprintf("%s/verify-account?id=%s&token=%s", config.EnvVars.FRONT_URL, strconv.Itoa(int(id)), token)

	err = tmpl.Execute(&bodyContentBuffer, struct {
		Name      string
		UrlSignup string
	}{
		Name:      name,
		UrlSignup: url,
	})
	if err != nil {
		fmt.Println(err)
	}

	SendEmail(to, subject, bodyContentBuffer.String())

}
