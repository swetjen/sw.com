package db

// collection names

const (
	Contacts = "contacts"
)

type (
	// ContactForm Describes a contact request
	ContactForm struct {
		Name      string
		Email     string
		Message   string
		IP        string
		TSCreated int64
	}
)
