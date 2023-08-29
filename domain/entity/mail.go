package entity

type Mail struct {
	Subject string            `json:"subject"`
	To      string            `json:"to"`
	Type    string            `json:"type"`
	Content map[string]string `json:"content"`
}

type Mails struct {
	MailContent []MailContent `json:"mailContent"`
}

type MailContent struct {
	SenderName string            `json:"senderName"`
	Subject    string            `json:"subject"`
	To         []string          `json:"to"`
	Cc         []string          `json:"cc"`
	Bcc        []string          `json:"bcc"`
	Type       string            `json:"type"`
	Content    map[string]string `json:"content"`
	SendAt     int64             `json:"sendAt"`
}
