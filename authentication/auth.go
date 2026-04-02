package authentication

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Email  string `json:"email"`
	UserID uint   `json:"userid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string, role string) (string, error) {
	claims := MyCustomClaims{
		email,
		userID,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	SigningKeyString := os.Getenv("JWT_SECRET")
	mySigningKey := []byte(SigningKeyString)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signature, err := token.SignedString(mySigningKey)
	SignedString := string(signature)
	return SignedString, err
}
