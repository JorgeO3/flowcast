// Package services contains the utility functions for the auth module.
package service

import (
	"gitlab.com/JorgeO3/flowcast/pkg/smtp"
)

// Mailer is an interface for sending emails.
type Mailer interface {
	SendConfirmationEmail(*MailerConfig) error
}

// MailerService is a service that sends emails.
type MailerService struct {
	smtpClient *smtp.Client
}

// MailerConfig represents the configuration required to send an email.
type MailerConfig struct {
	data         map[string]string
	email        string
	templateHTML string
	templateName string
}

// NewMailerConfig creates a new instance of MailerConfig.
func NewMailerConfig(data map[string]string, email string, templateHTML string, templateName string) *MailerConfig {
	return &MailerConfig{
		data:         data,
		email:        email,
		templateHTML: templateHTML,
		templateName: templateName,
	}
}

// SendConfirmationEmail sends a confirmation email.
func (m *MailerService) SendConfirmationEmail(cfg *MailerConfig) error {
	err := m.smtpClient.SendEmail(cfg.data, cfg.email, cfg.templateHTML, cfg.templateName)
	return err
}

// NewMailerService creates a new instance of MailerService.
func NewMailerService(smtpClient *smtp.Client) Mailer {
	return &MailerService{smtpClient: smtpClient}
}
