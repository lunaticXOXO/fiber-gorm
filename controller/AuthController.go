package controller

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("secret_key")
var tokenName = "auth_token"

type Claims struct{
	Username string `json:"username"`
	Type int `json:"type"`
	jwt.StandardClaims
}

func ResetToken(c *fiber.Ctx){
	c.Cookie(&fiber.Cookie{
		Name : tokenName,
		Value : "",
		Expires: time.Now(),
		Secure: false,
		HTTPOnly: true,
	})
}


func GenerateToken(c *fiber.Ctx,user model.Users) (error,string) {
	tokenExpiryTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: user.Username,
		Type: user.Type,
		StandardClaims:	jwt.StandardClaims{
			ExpiresAt : tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err,""
	}
	
	c.Cookie(&fiber.Cookie{
		Name:     tokenName,
		Value:    tokenString,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HTTPOnly: true,
	})

	return nil,tokenString
}

// Untuk authorization
func ValidateCookies(c *fiber.Ctx) (bool, string, int) {
	if cookie := c.Cookies(tokenName); cookie != ""{
		accessClaims := &Claims{}
		accessToken := cookie
		parseToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err == nil && parseToken.Valid {
			return true,accessClaims.Username, accessClaims.Type
		}
	}	

	return false, "", -1
}


func ValidateUserToken(c *fiber.Ctx, accessType int) bool{
	isAccess,username,usertype := ValidateCookies(c)
	fmt.Print("username : ",username,"usertype : ",usertype,"access : ",isAccess)

	if isAccess{
		userValid := usertype == accessType
		if userValid{
			return true
		}
	}
	return false
}

func Authenticate(next fiber.Handler, accessType int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isValidToken := ValidateUserToken(c,accessType)
		if !isValidToken {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "unauthorized access",
			})
		} else {
			return next(c)
		}
		return nil
	}
}

