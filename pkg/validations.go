package pkg

import (
	"errors"
	"regexp"
)

var commonPasswords = []string{
	"password",
	"123456",
	"123456789",
	"12345678",
	"12345",
	"1234567",
	"1234567890",
	"qwerty",
	"abc123",
	"password1",
}

func ValidateName(name string) error {
	regex := regexp.MustCompile(`^[A-Za-z]{4,}$`)
	if !regex.MatchString(name) {
		return errors.New("name must contain only letters and be longer than 3 characters")
	}
	return nil
}

func ValidateAge(age int) error {
	if age <= 18 {
		return errors.New("age must be greater than 18")
	}
	return nil
}

func ValidatePhone(phone string) error {
	regex := regexp.MustCompile(`^\+998\d{9}$`)
	if !regex.MatchString(phone) {
		return errors.New("phone number must be in the format +998XXXXXXXXX")
	}
	return nil
}

func ValidateUsername(username string) error {

    if len(username) < 5 {
        return errors.New("username must be at least 5 characters long")
    }

    if len(username) > 30 {
        return errors.New("username must be at most 30 characters long")
    }

    var validFirstChar = regexp.MustCompile(`^[a-zA-Z]`)
    if !validFirstChar.MatchString(username) {
        return errors.New("username must start with a letter")
    }

    var validUsername = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
    if !validUsername.MatchString(username) {
        return errors.New("username can only contain letters, numbers, and underscores")
    }
    return nil
}

func ValidateStatus(status string) error {
	allowedStatuses := map[string]bool{
		"waiting":  true,
		"active":   true,
		"pure":     true,
		"done":     true,
		"returned": true,
		"deleted":  true,
	}

	if _, ok := allowedStatuses[status]; !ok {
		return errors.New("invalid status")
	}
	return nil
}

func isCommonPassword(password string) bool {
	for _, commonPassword := range commonPasswords {
		if password == commonPassword {
			return true
		}
	}
	return false
}

func isRepeatedCharacter(password string) bool {
	if len(password) == 0 {
		return false
	}

	firstChar := password[0]
	for i := 1; i < len(password); i++ {
		if password[i] != firstChar {
			return false
		}
	}
	return true
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	if isCommonPassword(password) {
		return errors.New("password is too common")
	}

	if isRepeatedCharacter(password) {
		return errors.New("password cannot consist of the same symbol repeated")
	}

	return nil
}
