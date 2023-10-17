package config

import (
	"fmt"
	"os"
)

type Config struct {
	mailCfg MailConfig
}

type MailConfig struct {
	SMTPServer string
	SMTPPort   string
	Username   string
	Password   string
}

func New() (*Service, error) {
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	if smtpServer == "" || smtpPort == "" || username == "" || password == "" {
		return nil, fmt.Errorf("not all variables provided")
	}

	mailcfg := MailConfig{
		SMTPServer: smtpServer,
		SMTPPort:   smtpPort,
		Username:   username,
		Password:   password,
	}

	return &Service{cfg: Config{mailCfg: mailcfg}}, nil
}

type Service struct {
	cfg Config
}

func (s *Service) MailConfig() MailConfig {
	return s.cfg.mailCfg
}
