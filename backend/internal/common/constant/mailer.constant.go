package constant

type MailerTemplate string
type MailerSubject string

const (
	MailerUploadSuccessTemplate  MailerTemplate = "upload-success.html"
	MailerForgotPasswordTemplate MailerTemplate = "send-token-forgot-password.html"
	MailerCreateAccountAgent     MailerTemplate = "create-account-agent.html"
	MailerVerifyEmailAgent       MailerTemplate = "verify-email.html"
)

const (
	UploadSuccessSubject   MailerSubject = "Boilerplate: Sukses Mengunggah File | Success Upload file"
	ForgotPasswordTemplate MailerSubject = "Forgot Password: Request Change Password"
	ForgotVerifyEmailAgent MailerSubject = "Verify Email: Create Agent Account"
)
