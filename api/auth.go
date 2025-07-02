package api

import (
  "time"
  "net/http"
  "errors"
  "strings"
  "context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
  "github.com/gin-gonic/gin"

  "main/model"
)

func login(c *gin.Context) {
  var data model.User

  if datas != nil {
    data = datas.GetUser()
  }

  token, err := generateJWT(data.Id)

  if err != nil {
    c.JSON(500, "Error while generating JWT token : " + err.Error())
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

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := extractUserIDFromJWT(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractUserIDFromJWT(r *http.Request) (uuid.UUID, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return uuid.Nil, errors.New("missing auth header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return uuid.Nil, errors.New("invalid auth header format")
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
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

