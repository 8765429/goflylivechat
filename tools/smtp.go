package tools

import (
	"crypto/tls"
	"encoding/base64"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"io"
	"strings"
)

func SendSmtp(server string, from string, password string, to []string, subject string, body string) error {
	auth := sasl.NewPlainClient("", from, password)
	subjectBase := base64.StdEncoding.EncodeToString([]byte(subject))
	msgStr := "From: " + from + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: =?UTF-8?B?" + subjectBase + "?=\r\n" +
		"Content-Type: text/html; charset=UTF-8" +
		"\r\n\r\n" +
		body + "\r\n"

	if strings.HasSuffix(server, ":465") {
		host := strings.Split(server, ":")[0]
		tlsConfig := &tls.Config{ServerName: host}
		c, err := smtp.DialTLS(server, tlsConfig)
		if err != nil {
			return err
		}
		defer c.Close()

		if err = c.Auth(auth); err != nil {
			return err
		}
		if err = c.Mail(from, nil); err != nil {
			return err
		}
		for _, addr := range to {
			if err = c.Rcpt(addr); err != nil {
				return err
			}
		}
		w, err := c.Data()
		if err != nil {
			return err
		}
		_, err = io.Copy(w, strings.NewReader(msgStr))
		if err != nil {
			return err
		}
		if err = w.Close(); err != nil {
			return err
		}
		return c.Quit()
	}

	return smtp.SendMail(server, auth, from, to, strings.NewReader(msgStr))
}
