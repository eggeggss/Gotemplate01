package controller

import (
	"fmt"
	models "hw/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwt secret key
var jwtSecret = []byte("secret")

func AuthRequired(c *gin.Context) {

	auth := c.GetHeader("Authorization")

	token := strings.Split(auth, "Bearer ")[1]

	tokenClaims, err := jwt.ParseWithClaims(token, &models.Claims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtSecret, nil
		})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*models.Claims); ok && tokenClaims.Valid {
		fmt.Println("account:", claims.Account)
		fmt.Println("role:", claims.Role)
		c.Set("account", claims.Account)
		c.Set("role", claims.Role)
		c.Next()
	} else {
		c.Abort()
		return
	}
}

func LoginUser(c *gin.Context) {

	var u models.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if u.Name == "admin" {

		now := time.Now()
		jwtId := u.Name + strconv.FormatInt(now.Unix(), 10)
		role := "Member"

		claims := models.Claims{
			Account: u.Name,
			Role:    role,
			StandardClaims: jwt.StandardClaims{
				Audience:  u.Name,
				ExpiresAt: now.Add(60 * time.Second).Unix(),
				Id:        jwtId,
				IssuedAt:  now.Unix(),
				Issuer:    "ginJWT",
				NotBefore: now.Add(1 * time.Second).Unix(),
				Subject:   u.Name,
			},
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusForbidden)
	}

}
