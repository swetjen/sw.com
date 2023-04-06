package db

// ContactForm Describes a contact request
type ContactForm struct {
	Name      string
	Email     string
	Phone     string
	Message   string
	IP        string
	TSCreated int64
}
