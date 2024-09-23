package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Fadhelbulloh/DBO-Test-Case/params"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string, exp int64) (string, error) {

	expiredTimeString := os.Getenv("JWT_HOUR_EXPIRATION")
	expiredTime, err := strconv.Atoi(expiredTimeString)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return os.Getenv("JWT_SECRET"), nil
	})
}

// ReturnErrorTemplate returns error template
func ReturnErrorTemplate(err error, message string) map[string]interface{} {
	if err != nil {
		log.Println(err)
	} else {
		log.Println("error :" + message)
	}
	return map[string]interface{}{
		"status": "failed",
		"data":   message,
	}
}

// ReturnSuccessTemplate returns success template
func ReturnSuccessTemplate(data interface{}, pagination ...params.Pagination) map[string]interface{} {
	return map[string]interface{}{
		"status":     "success",
		"data":       data,
		"pagination": pagination,
	}
}
