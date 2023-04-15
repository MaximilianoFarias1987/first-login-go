package helpers

import (
	"login/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["userName"] = user.UserName
	claims["name"] = user.Name
	claims["lastName"] = user.LastName
	claims["userEmail"] = user.Email
	claims["userAddress"] = user.Address
	claims["userPhoneNumber"] = user.PhoneNumber
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Vencimiento en 24 horas

	tokenString, err := token.SignedString([]byte("clave-secreta"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
