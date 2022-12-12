package middleware

import (
	"os"
	"time"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"capstone-project/entity"
)

func CreateToken(username string, email string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func CreateRefreshToken(username string, email string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractTokenUsername(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["username"].(string)
}

// func VerifyToken(refreshToken string) (*jwt.PayloadBinding, error) {
// 	keyFunc := func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 		if !ok {
// 			return nil, errors.New("Invalid token")
// 		}
// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	}

// 	jwtToken, err := jwt.ParseWithClaims(refreshToken, &entity.PayloadBinding{}, keyFunc)
// 	if err != nil {
// 		verr, ok := err.(*jwt.ValidationError)
// 		if ok && errors.Is(verr.Inner, errors.New("Token expired")) {
// 			return nil, errors.New("Token expired")
// 		}
// 		return nil, errors.New("Invalid token")
// 	}

// 	payload, ok := jwtToken.Claims.(*entity.PayloadBinding)
// 	if !ok {
// 		return nil, errors.New("Invalid token")
// 	}

// 	return payload, nil
// }

func ValidateToken(refreshToken entity.TokenBinding) (*entity.Users, error) {

	
	// var refToken entity.TokenBinding
	// refToken = refreshToken;
	fmt.Println("ini refresh token ", string(refreshToken.RefreshToken));
	token, err := jwt.Parse(refreshToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	user := &entity.Users{}
	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("this is payload ", payload);
	if ok && token.Valid {
		user.Username = payload["username"].(string)
		user.Email = payload["email"].(string)

		return user, nil
	}

	return nil, errors.New("invalid token")
}