package helper

import (
	"api/model"
	"errors"
	"fmt"
	"strings"
	"time"
	//"os"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
//"strconv"
)

var privateKey = []byte(viper.GetString("app.privatkey"))

func GenerateJWT(user model.User) (string, error) {
	//tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		//"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(c *fiber.Ctx) error {
	token, err := getToken(c)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func CurrentUser(c *fiber.Ctx) (model.User, error) {
	err := ValidateJWT(c)
	if err != nil {
		return model.User{}, err
	}

	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := model.FindUserById(userId)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func getToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(c *fiber.Ctx) string {
	token := c.Cookies("token")
	if token != "" {
		return token
	}

	head := c.GetReqHeaders()
	bearerToken := head["Authorization"]
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
