package mail

import (
	"backend_template/src/core/domain/crawler"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/mail.v2"
)

func SendListLinksEmail(links []crawler.VisitedLinks, email string) error {
	title := fmt.Sprintf("Lista de links encontrados")
	list := ""
	for _, link := range links {
		list = fmt.Sprintf(`%v <p> Website: %s; URL: %s; Data: %v</p>
									 </br>`, list, link.Website(), link.Url(), link.CheckedDate())
	}
	body := fmt.Sprintf(`
		<h2>Lista de links encontrado</h2>
		</br>
		</br>
		</br>
		%s
	`, list)
	return send(title, body, []string{email})
}

func send(title, body string, emails []string) error {
	message := mail.NewMessage()
	message.SetHeader("From", getFromName())
	message.SetHeader("To", strings.Join(emails, ","))
	message.SetHeader("Subject", title)
	message.SetHeader("Sender", getFromAddress())
	message.SetBody("text/html", body)
	dialer, err := newDialer()
	if err != nil {
		return err
	}
	return dialer.DialAndSend(message)
}

func newDialer() (*mail.Dialer, error) {
	serviceHost := os.Getenv("MAIL_SMTP_HOST")
	servicePort := os.Getenv("MAIL_SMTP_PORT")
	fromEmail := getFromAddress()
	fromPassword := os.Getenv("MAIL_FROM_PASSWORD")
	if serviceHost == "" {
		return nil, errors.New("you need to define the mail service host!")
	} else if servicePort == "" {
		return nil, errors.New("you need to define the mail service port!")
	} else if fromEmail == "" {
		return nil, errors.New("you need to define the responsible mail email address!")
	} else if fromPassword == "" {
		return nil, errors.New("you need to define the responsible email password!")
	} else if _, err := strconv.Atoi(servicePort); err != nil {
		return nil, errors.New("the mail service port must be an integer!")
	} else if getFromName() == "" {
		return nil, errors.New("you need to define the email sender name!")
	}
	intServicePort, _ := strconv.Atoi(servicePort)
	return mail.NewDialer(serviceHost, intServicePort, fromEmail, fromPassword), nil
}

func getFromAddress() string {
	return os.Getenv("MAIL_FROM_ADDRESS")
}

func getFromName() string {
	return os.Getenv("MAIL_FROM_NAME")
}
