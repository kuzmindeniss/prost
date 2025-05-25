package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const JWTExp = time.Hour * 24

func CreateToken(userUUID uuid.UUID) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userUUID.String(),
		"exp": time.Now().Add(JWTExp).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func ReadToken(tokenString string) (userUUID uuid.UUID, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("ошибка при парсинге токена: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, errors.New("не удалость прочитать токен")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return uuid.UUID{}, errors.New("время действия токена истекло")
	}

	userIdString := claims["sub"].(string)
	userIdUUID, err := uuid.Parse(userIdString)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("ошибка при парсинге id пользователя в uuid: %v", err)
	}

	return userIdUUID, nil
}
