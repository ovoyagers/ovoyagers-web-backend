package inviteservice

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/petmeds24/backend/config"
	log "github.com/sirupsen/logrus"
)

//go:embed templates/inviteuser.html
var emailTemplate embed.FS

func (is *InviteService) SendInviteEmail(email string, inviteId string) error {
	// Load the config from the environment
	cfg, err := config.LoadConfig() // Load the config from the environment
	if err != nil {
		return err
	}

	// Determine the sender address and password
	const templatePath = "templates/inviteuser.html" // Path to the email template
	user := cfg.EMAIL_USER                           // SMTP username
	password := cfg.EMAIL_PASSWORD                   // SMTP password
	host := cfg.EMAIL_HOST                           // SMTP host
	port := cfg.EMAIL_PORT                           // SMTP port
	addr := fmt.Sprintf("%s:%s", host, port)         // SMTP address

	// Read the email template from the embedded filesystem
	emailContent, err := emailTemplate.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// Parse the email template
	t, err := template.New("inviteuser").Parse(string(emailContent))
	if err != nil {
		return err
	}

	// Prepare the data for the email template
	data := struct {
		Username       string
		InvitationLink string
	}{
		Username:       email,
		InvitationLink: inviteId,
	}

	// Generate the email body
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Pet Store - OTP Verification Code\n%s", mimeHeaders)))

	// Execute the email template with the data
	err = t.Execute(&body, data)

	// Check if there was an error
	if err != nil {
		return err
	}

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", user, password, host)

	// Send the email
	if err := smtp.SendMail(addr, auth, user, []string{email}, body.Bytes()); err != nil {
		return err
	}

	// Log that the OTP was sent
	log.Printf("OTP sent to %s\n", email)

	// Return nil if the email was sent successfully
	return nil
}
