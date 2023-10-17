package mail

import (
	"crypto/tls"
	"io"
	"net/smtp"
	"path/filepath"

	"github.com/iliyasali2107/archiver/internal/config"
	"github.com/iliyasali2107/archiver/internal/dto"
	"github.com/jordan-wright/email"
)

type MailSvc struct {
	cfg MailConfigGetter
}

type MailConfigGetter interface {
	MailConfig() config.MailConfig
}

func NewMailSvc(config MailConfigGetter) *MailSvc {
	return &MailSvc{
		cfg: config,
	}
}

func (mss *MailSvc) SendMail(req dto.SendMailRequest) error {
	fileHeader := req.FileHeader
	emails := req.ReceiverEmails
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	smtpServer := mss.cfg.MailConfig().SMTPServer
	smtpPort := mss.cfg.MailConfig().SMTPPort
	username := mss.cfg.MailConfig().Username
	password := mss.cfg.MailConfig().Password

	auth := smtp.PlainAuth("", username, password, smtpServer)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpServer,
	}

	client, err := smtp.Dial(smtpServer + ":" + smtpPort)
	if err != nil {
		return err
	}
	defer client.Quit()

	err = client.StartTLS(tlsConfig)
	if err != nil {
		return err
	}

	err = client.Auth(auth)
	if err != nil {
		return err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	e := email.NewEmail()
	e.From = "<ilyasaliev1@gmail.com>"
	e.To = emails
	e.Subject = "Subject"
	e.Text = []byte("Text")
	_, fileName := filepath.Split(fileHeader.Filename)
	at, err := e.Attach(file, fileName, "application/octet-stream")
	if err != nil {
		return err
	}

	at.Header = fileHeader.Header
	at.Content = content
	err = e.Send(smtpServer+":"+smtpPort, auth)
	if err != nil {
		return err
	}

	return nil
}
