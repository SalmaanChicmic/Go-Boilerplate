package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
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

func DecodeToken(tokenString string) (*jwt.MapClaims, error) {

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You need to provide the key to validate the signature
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	// Check for errors
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		fmt.Println("Token is not valid")
		return nil, err
	}

	// Access claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error accessing claims")
		return nil, err
	}

	return &claims, nil
}

func HashPassword(password string) (*string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(bs)
	return &hashedPassword, nil
}

func CheckPasswordHash(password, hash string) bool {

	fmt.Println("inside password check ")
	fmt.Println("password hash", password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
