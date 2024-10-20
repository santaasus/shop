package domain

type Config struct {
	SMTP SMTP `json:"SMTP"`
}

type SMTP struct {
	Source   string `json:"smtp_source"`
	MailFrom string `json:"mail_from"`
}

type NotificationInfo struct {
	OrderId   int
	UserEmail string
}
