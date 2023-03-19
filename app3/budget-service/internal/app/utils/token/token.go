package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user_id string) (string, error) {
	token_lifespan := 3600
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(token_lifespan)).Unix(),
		"iat": time.Now().Unix(),
		"sub": user_id,
	}).SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractTokenID(c *gin.Context) (string, error) {
	tokenString := ExtractToken(c)
	tokenClaims, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims := tokenClaims.Claims.(*jwt.MapClaims)
	user_id := (*claims)["sub"].(string)
	return user_id, nil
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
