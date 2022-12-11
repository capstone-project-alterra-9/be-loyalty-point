package helper

func ValidateLength(value string, min int, max int) bool {
	if len(value) < min || len(value) > max {
		return false
	}
	return true
}

func ValidateAlphanumeric(value string) bool {
	for _, v := range value {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func ValidateEmail(value string) (string, error) {
	var (
		subEmail string
		valid    bool = false
	)
	for _, v := range value {
		if v != '@' {
			subEmail += string(v)
		} else {
			valid = true
			break
		}
	}
	if !valid {
		return "", ErrEmailValid
	}
	if len(subEmail) < 8 {
		return "", ErrEmailLength
	}

	return subEmail, nil
}
