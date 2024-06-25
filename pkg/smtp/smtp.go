// Package smtp provides a simple interface for sending HTML emails using the SMTP protocol.
//
// Example usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "log"
//	    "os"
//	    "your_project/smtp" // Actualiza esto con el nombre correcto de tu módulo
//
//	    "github.com/joho/godotenv"
//	)
//
//	func main() {
//	    // Load environment variables from .env file
//	    if err := godotenv.Load(); err != nil {
//	        log.Fatalf("Error loading .env file: %v", err)
//	    }
//
//	    // Load configuration from environment variables
//	    host := os.Getenv("SMTP_HOST")
//	    port := os.Getenv("SMTP_PORT")
//	    from := os.Getenv("SMTP_FROM")
//	    password := os.Getenv("SMTP_PASSWORD")
//	    templatePath := os.Getenv("TEMPLATE_PATH")
//
//	    if host == "" || port == "" || from == "" || password == "" || templatePath == "" {
//	        log.Fatal("Missing required environment variables")
//	    }
//
//	    // Create SMTP config and client
//	    smtpConfig := smtp.NewConfig(host, port, from, password)
//	    smtpClient := smtp.NewSMTPClient(smtpConfig)
//
//	    // Prepare email data
//	    data := map[string]string{
//	        "Name": "Jorge",
//	        "Age":  "24",
//	    }
//
//	    // Read the HTML template from file
//	    templateHTML, err := os.ReadFile(templatePath)
//	    if err != nil {
//	        log.Fatalf("Error reading template file: %v", err)
//	    }
//
//	    // Send email
//	    err = smtpClient.SendEmail(data, "joheosmo@gmail.com", string(templateHTML), "email_template")
//	    if err != nil {
//	        fmt.Println("Error sending email:", err)
//	    } else {
//	        fmt.Println("Email sent successfully")
//	    }
//	}
package smtp

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

// Config represents the configuration required to send an email.
type Config struct {
	Host     string // SMTP server host
	Port     string // SMTP server port
	From     string // Sender email address
	Password string // Sender email password
}

// NewConfig creates a new instance of Config.
//
// Parameters:
//   - host: SMTP server host.
//   - port: SMTP server port.
//   - from: Sender email address.
//   - password: Sender email password.
//
// Returns:
//   - A pointer to the new Config instance.
func NewConfig(host, port, from, password string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		From:     from,
		Password: password,
	}
}

// Client represents the SMTP service.
//
// Example usage:
//
//	cfg := smtp.NewConfig("smtp.example.com", "587", "your-email@example.com", "your-password")
//	client := smtp.NewSMTPClient(cfg)
//	data := map[string]string{
//	    "Name": "Jorge",
//	    "Age":  "24",
//	}
//	templateHTML := `
//	    <!DOCTYPE html>
//	    <html lang="en">
//	    <head>
//	        <meta charset="UTF-8">
//	        <meta http-equiv="X-UA-Compatible" content="IE=edge">
//	        <meta name="viewport" content="width=device-width, initial-scale=1.0">
//	        <title>Email Template</title>
//	    </head>
//	    <body>
//	        <h1>Hello {{.Name}}</h1>
//	        <p>Your age is {{.Age}}</p>
//	    </body>
//	    </html>
//	`
//	err := client.SendEmail(data, "recipient@example.com", templateHTML, "email_template")
//	if err != nil {
//	    fmt.Println("Error sending email:", err)
//	}
type Client struct {
	config *Config // SMTP configuration
}

// NewSMTPClient creates a new instance of SMTPClient.
//
// Parameters:
//   - cfg: Pointer to a Config instance.
//
// Returns:
//   - A pointer to the new SMTPClient instance.
func NewSMTPClient(cfg *Config) *Client {
	return &Client{config: cfg}
}

// SendEmail sends an email using the specified data and template.
//
// Parameters:
//   - data: Data to be injected into the email template.
//   - email: Recipient email address.
//   - tmpl: HTML email template as a string.
//   - tmplName: Name of the email template.
//
// Returns:
//   - An error if the email could not be sent, otherwise nil.
//
// Example usage:
//
//	data := map[string]string{
//	    "Name": "Jorge",
//	    "Age":  "24",
//	}
//	templateHTML := `
//	    <!DOCTYPE html>
//	    <html lang="en">
//	    <head>
//	        <meta charset="UTF-8">
//	        <meta http-equiv="X-UA-Compatible" content="IE=edge">
//	        <meta name="viewport" content="width=device-width, initial-scale=1.0">
//	        <title>Email Template</title>
//	    </head>
//	    <body>
//	        <h1>Hello {{.Name}}</h1>
//	        <p>Your age is {{.Age}}</p>
//	    </body>
//	    </html>
//	`
//	err := client.SendEmail(data, "recipient@example.com", templateHTML, "email_template")
//	if err != nil {
//	    fmt.Println("Error sending email:", err)
//	}
func (s *Client) SendEmail(data any, email, tmpl, tmplName string) error {
	t, err := template.New(tmplName).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Simplified message creation
	msg := fmt.Sprintf(
		"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: Test Email\r\n"+
			"\r\n%s",
		s.config.From, email, body.String())

	auth := smtp.PlainAuth("", s.config.From, s.config.Password, s.config.Host)
	to := []string{email}
	err = smtp.SendMail(s.config.Host+":"+s.config.Port, auth, s.config.From, to, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
