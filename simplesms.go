package simplesms

import (
	"fmt"
	"net/http"
	"net/smtp"
)

const (
	TMO string = "TMO"
	ATT string = "ATT"
	VRZ string = "VRZ"
)

func CheckProvider(prov string) error {
	if prov != TMO && prov != ATT && prov != VRZ {
		return fmt.Errorf("invalid provider specified. options are 'TMO', 'ATT', or 'VRZ'")
	}
	return nil
}

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

func (c *Client) Send(phoneNum int, prov string, subject string, message string) error {
	var recipient string
	switch prov {
	case TMO:
		recipient = fmt.Sprintf("%d@tmomail.net", phoneNum)
	case ATT:
		recipient = fmt.Sprintf("%d@txt.att.net", phoneNum)
	case VRZ:
		recipient = fmt.Sprintf("%d@vtext.com", phoneNum)
	default:
		return fmt.Errorf("invalid provider specified: %s", prov)
	}

	email := "From: " + c.username + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	err := smtp.SendMail(c.smtpHost+c.smtpPort,
		smtp.PlainAuth("", c.username, c.password, c.smtpHost),
		c.username, []string{recipient}, []byte(email))

	if err != nil {
		return fmt.Errorf("unable to send message: %s", err)
	}
	return nil

}
