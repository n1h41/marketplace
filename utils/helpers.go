package utils

import "regexp"

func ValidateEmail(email string) bool {
	isValid, err := regexp.MatchString(`^\w+@\w+\.\w+$`, email)
	if err != nil {
		return false
	}
	return isValid
}
