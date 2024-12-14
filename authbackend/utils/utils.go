package utils

func IsValidPhone(phone string) bool {
	// Check if the input is empty.
	if len(phone) == 0 {
		return false
	}

	if phone[0] == '+' {
		phone = phone[1:]
	}

	// Ensure the phone number has 10-13 digits.
	if len(phone) < 10 || len(phone) > 13 {
		return false
	}

	for _, char := range phone {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}
