package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"main/model"
)

func register(c *gin.Context) {
	var newUser model.User
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		msg := gin.H{
			"error": "Could not read request body " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	var d map[string]string
	if err := json.Unmarshal(jsonData, &d); err != nil {
		msg := gin.H{
			"error": "invalid JSON format " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	var user, password string = d["username"], d["password"]

	if user == "" || password == "" {
		msg := gin.H{
			"error": "invalid body fields, need username and password",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	newUser.Init(user, password)
	if datas != nil {
		fmt.Println("NEW USER -> UUID : " + newUser.Id.String() + " USERNAME : " + newUser.Username + " PASSWORD : " + newUser.Password)
	}

	c.JSON(http.StatusOK, "New user created")
}

func login(c *gin.Context) {
	var data model.User
	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		msg := gin.H{
			"error": "Could not read request body " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	var d map[string]string
	if err := json.Unmarshal(jsonData, &d); err != nil {
		msg := gin.H{
			"error": "invalid JSON format " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	var user, password string = d["username"], d["password"]

	if user == "" || password == "" {
		msg := gin.H{
			"error": "invalid body fields, need username and password",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	if datas != nil {
		var u = datas.GetUser()
		if u.Username == user && u.Password == password {
			data = u
		}
	}

	if (data == model.User{}) {
		msg := gin.H{
			"error": "user not found, wrong credentials",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, msg)
		return
	}

	token, err := generateJWT(data.Id)

	if err != nil {
		c.AbortWithStatusJSON(500, "Error while generating JWT token : "+err.Error())
	}

	c.JSON(http.StatusOK, token)
}

func generateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := extractUserIDFromJWT(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		}

		c.Set("UserId", userID)
		c.Next()
	}
}

func extractUserIDFromJWT(c *gin.Context) (uuid.UUID, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return uuid.Nil, errors.New("missing auth header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return uuid.Nil, errors.New("invalid auth header format")
	}

	tokenStr := parts[1]

	fmt.Println(tokenStr)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.Nil, errors.New("missing sub claim")
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, errors.New("invalid uuid in token")
	}

	return userID, nil
}

func getUserIdFromContext(c *gin.Context) (any, error) {
	userId, exist := c.Get("UserId")

	if !exist {
		return nil, errors.New("user id not found in context")
	}

	return userId, nil
}
