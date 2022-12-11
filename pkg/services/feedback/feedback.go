package feedback

import "gopkg.in/gomail.v2"

func Feedback(email string, text string) {
	msg := setAttributes(email, text)
	sendEmail(msg)
}

func setAttributes(email string, text string) *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", email) //"callbackuser25@gmail.com"
	msg.SetHeader("To", "tournamentmaker25@gmail.com")
	msg.SetHeader("Subject", email)
	msg.SetBody("text/html", text)

	return msg
}

func sendEmail(msg *gomail.Message) {
	n := gomail.NewDialer("smtp.gmail.com", 587, "callbackuser25@gmail.com", "fwdgxuxsfduozwre")

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
