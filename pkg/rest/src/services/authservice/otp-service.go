package authservice

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	log "github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

//go:embed templates/otp.html
var emailTemplate embed.FS

func getTwillioClient() (*twilio.RestClient, string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, "", err
	}
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TWILLIO_ACCOUNT_SID,
		Password: cfg.TWILLIO_ACCOUNT_PASSWORD,
	})
	return client, cfg.TWILLIO_ACCOUNT_SERVICE_SID, nil
}

func (authService *AuthService) SendOTP(phone string) (string, error) {

	client, serviceID, err := getTwillioClient()
	if err != nil {
		return "", err
	}

	params := &openapi.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func (authService *AuthService) VerifyOTP(otp authmodel.OTP) (string, error) {
	client, serviceID, err := getTwillioClient()
	if err != nil {
		return "", err
	}

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(otp.Phone)
	params.SetCode(otp.Code)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)
	if err != nil {
		return "", err
	}
	return *resp.Status, nil
}

// SendOTPViaEmail sends an OTP via email to the provided email address.
// It loads the email configuration from the environment, reads an email template from the embedded filesystem,
// populates the template with the verification code, and sends the email using the configured SMTP server.
//
// Parameters:
// - to: the email address to send the OTP to
// - code: the verification code to include in the email
//
// Returns:
// - error: an error if there was a problem sending the email
func (authService *AuthService) SendOTPViaEmail(to string, code string) error {
	// Load the config from the environment
	cfg, err := config.LoadConfig() // Load the config from the environment
	if err != nil {
		return err
	}

	// Determine the sender address and password
	const templatePath = "templates/otp.html" // Path to the email template
	user := cfg.EMAIL_USER                    // SMTP username
	password := cfg.EMAIL_PASSWORD            // SMTP password
	host := cfg.EMAIL_HOST                    // SMTP host
	port := cfg.EMAIL_PORT                    // SMTP port
	addr := fmt.Sprintf("%s:%s", host, port)  // SMTP address

	// Read the email template from the embedded filesystem
	emailContent, err := emailTemplate.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// Parse the email template
	t, err := template.New("otp").Parse(string(emailContent))
	if err != nil {
		return err
	}

	// Prepare the data for the email template
	data := struct {
		Username         string
		VerificationCode string
	}{
		Username:         to,
		VerificationCode: code,
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
	if err := smtp.SendMail(addr, auth, user, []string{to}, body.Bytes()); err != nil {
		return err
	}

	// Log that the OTP was sent
	log.Printf("OTP sent to %s\n", to)

	// Return nil if the email was sent successfully
	return nil
}

func (authService *AuthService) VerifyOTPViaEmail(request authmodel.VerifyEmailRequest) (map[string]interface{}, error) {
	otpMap, err := authService.authDao.GetOTP(request.Email)
	if err != nil {
		return nil, err
	}
	// Compare the provided OTP with the one stored in the database
	if otpMap["otp"] != request.Code {
		return nil, fmt.Errorf("invalid otp")
	}
	result, err := authService.authDao.VerifyOTP(request.Email)
	return result, err
}

func (authService *AuthService) UpdateOTP(email string, otp string) error {
	return authService.authDao.UpdateOTP(email, otp)
}
