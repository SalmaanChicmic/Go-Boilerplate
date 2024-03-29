package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func IsPassValid(password string) (bool, error) {

	if len(password) < 8 {
		return false, errors.New("password is too short")

	}
	hasUpperCase := false
	hasLowerCase := false
	hasNumbers := false
	hasSpecial := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpperCase = true
		} else if char >= 'a' && char <= 'z' {
			hasLowerCase = true
		} else if char >= '0' && char <= '9' {
			hasNumbers = true
		} else if char >= '!' && char <= '/' {
			hasSpecial = true
		} else if char >= ':' && char <= '@' {
			hasSpecial = true
		}
	}

	if !hasUpperCase {
		return false, errors.New("password do not contain upperCase Character")
	}

	if !hasLowerCase {
		return false, errors.New("password do not contain LowerCase Character")
	}

	if !hasNumbers {
		return false, errors.New("password do not contain any numbers")
	}

	if !hasSpecial {
		return false, errors.New("password do not contain any special character")
	}
	return true, nil
}

func HashPassword(password string) (*string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(bs)
	return &hashedPassword, nil
}
