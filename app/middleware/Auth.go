package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func Auth(c *gin.Context){
	TokenString := c.Request.Header.Get("Authorization")
	token,err := jwt.Parse(TokenString, func(token *jwt.Token)(interface{},error){
		if jwt.GetSigningMethod("HS256") != token.Method{
			return nil, fmt.Errorf("unexpectrd signing method : %v", token.Header["alg"])
		}
		return []byte("secret"),nil

	})
	if err !=nil{
		result := gin.H{
			"message": "invalid token",
			"error":err.Error(),
		}
		c.JSON(http.StatusUnauthorized,result)
		c.Abort()
		return
	}
	fmt.Println("Token Verified")
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)

	var idAccount int

	err = mapstructure.Decode(claims["account_number"],&idAccount)
	if err != nil{
		result :=gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusUnauthorized,result)
		c.Abort()
	}
	fmt.Println(idAccount)
	c.Set("account_number", idAccount)}