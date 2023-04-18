package helper

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mujahxd/eventgenie/config"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {

	response := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return response

}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func GenerateToken(ID uint) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(config.TokenSecret))
	if err != nil {
		fmt.Println(err.Error())
		return signedToken, err
	}
	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return token, err
	}
	return token, nil
}

func DecodeToken(token *jwt.Token) (uint, error) {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		ID := data["id"].(float64)
		return uint(ID), nil
	}
	return 0, errors.New("failed to decode token")
}
