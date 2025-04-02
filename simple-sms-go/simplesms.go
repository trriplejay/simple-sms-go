package simplesms

import (
	"fmt"
	"net/http"
	"net/smtp"
)

type Provider int

const (
	TMO Provider = iota // 0
	ATT
	VRZ
)

type Client struct {
	client   *http.Client
	username string
	password string
	smtpHost string
	smtpPort string
}

func NewClient(username string, password string, smtpHost string, smtpPort string) *Client {
	return &Client{client: &http.Client{},
		username: username,
		password: password,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
	}
}

func (c *Client) Send(num int, prov Provider, subject string, message string) error {
	var recipient string
	switch prov {
	case TMO:
		recipient = fmt.Sprintf("%d@tmomail.net", num)
	case ATT:
		recipient = fmt.Sprintf("%d@txt.att.net", num)
	case VRZ:
		recipient = fmt.Sprintf("%d@vtext.com", num)
	default:
		return fmt.Errorf("Invalid provider specified: %d", prov)
	}

	email := "From: " + c.username + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	err := smtp.SendMail(c.smtpHost+c.smtpPort,
		smtp.PlainAuth("", c.username, c.password, c.smtpHost),
		c.username, []string{recipient}, []byte(email))

	if err != nil {
		return fmt.Errorf("Unable to send message: %s", err)
	}
	return nil

}
