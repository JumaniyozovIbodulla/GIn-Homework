package check

import (
	"errors"
	"regexp"
	"time"
)

func ValidateYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}

func ValidatePhone(phone string) error {
	isMatch, err := regexp.MatchString(`^\+998\d{2}\d{7}$`, phone)
	if isMatch && err == nil {
		return nil
	}
	return errors.New("phone number is not valid")
}

func ValidatePassword(password string) error {
	isMatch, err := regexp.MatchString(`^[a-zA-Z0-9]{8,}$`, password)

	if isMatch && err == nil {
		return nil
	}
	return errors.New("password is not valid")
}

func ValidateEmail(email string) error {
	isMatch, err := regexp.MatchString(`^.*@gmail\.com$`, email)

	if isMatch && err == nil {
		return nil
	}
	return errors.New("email is not valid")
}
