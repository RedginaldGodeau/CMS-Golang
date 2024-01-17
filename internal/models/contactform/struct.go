package contactform

type ContactFormModel struct {
	Username string
	Email    string
	Message  string
	ClientIp string
	SendAt   int64
}
