package token

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Generate JWT Token
func GenerateToken(claims model.Claims) (*string, error) {
	//create user claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))

	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

// Decode Token function
func DecodeToken(tokenString string) (*model.Claims, error) {
	claims := &model.Claims{}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if err != nil || !parsedToken.Valid {
		if claims.ExpiresAt != nil && (*claims.ExpiresAt).Before(time.Now()) {
			var userToBeLoggedOut model.User
			err := db.FindById(&userToBeLoggedOut, claims.Id, "user_id")
			if err != nil {
				return nil, fmt.Errorf("error finding user in db")
			}
			query := "UPDATE users SET is_active = false WHERE user_id = '" + claims.Id + "'"
			db.QueryExecutor(query, &userToBeLoggedOut)
			fmt.Println("user to be logged out ", userToBeLoggedOut)

			var userSessionToBeDeleted model.Session
			err = db.FindById(&userSessionToBeDeleted, claims.Id, "user_id")
			if err != nil {
				return nil, fmt.Errorf("error finding user in db")
			}

			db.DeleteRecord(&userSessionToBeDeleted, claims.Id, "user_id")

			return nil, fmt.Errorf("token has expired , please proceed to login")
		}
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
