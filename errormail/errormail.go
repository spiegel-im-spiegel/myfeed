package errormail

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/goark/errs"
	"github.com/spiegel-im-spiegel/myfeed/env"
)

const (
	subject = "error in " + env.ServiceName + " batch"
	message = "error in " + env.ServiceName + " batch.\n"
)

// SendErrorMail sends email in fixed message.
func SendErrorMail() error {
	cfg, err := env.EmailConfig()
	if err != nil {
		return errs.Wrap(err)
	}
	if cfg == nil {
		return nil
	}

	var auth smtp.Auth
	if cfg.Encrypt {
		auth = smtp.CRAMMD5Auth(cfg.Username, cfg.Password)
	} else {
		auth = smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Hostname)
	}
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port), auth, cfg.From, []string{cfg.To}, makeMsg(cfg.To)); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func makeMsg(to string) []byte {
	ss := []string{
		"To: " + to,
		"Subject: " + subject,
		"\n",
		message,
	}
	return []byte(strings.ReplaceAll(strings.Join(ss, "\n"), "\n", "\r\n"))
}
