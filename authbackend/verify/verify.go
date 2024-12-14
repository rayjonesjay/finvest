package verify

import "net/mail"

func Email(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NationalID(id string) bool {
	if id == "" || len(id) < 6 {
		return false
	}

	for _, cn := range id {
		if cn >= '0' && cn <= '9' {
			continue
		} else {
			return false
		}
	}

	return true
}
