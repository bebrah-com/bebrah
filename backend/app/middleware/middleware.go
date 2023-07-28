package middleware

import (
	"bebrah/app/db"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"bebrah/app/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pingcap/log"
	"go.uber.org/zap"
)

var SecretKey = []byte("secret")

const UserIdKey = "userId"

func GetUserIdFromGin(c *gin.Context) (uint64, error) {
	userId, exist := c.Get(UserIdKey)
	if !exist {
		return 0, errors.New("userId not found")
	}

	userIdStr, ok := userId.(string)
	if !ok {
		return 0, errors.New("userId is not string")
	}

	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, err
	}

	return uint64(userIdInt), nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		parts := strings.Split(tokenString, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			log.Info("invalid token", zap.String("token", tokenString))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		tokenString = parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					log.Info("That's not even a token")
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					log.Info("Timing is everything")
				} else {
					log.Info("Couldn't handle this token:", zap.Error(err))
				}
			} else {
				log.Info("Couldn't handle this token:", zap.Error(err))
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		}
		if token.Valid {
			log.Info("Get user from token", zap.Any("claims", claims))
			c.Set(UserIdKey, claims["jti"])

			userId, err := GetUserIdFromGin(c)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				return
			}

			// get user from db
			var user model.User
			db.Db().Where("id = ?", userId).First(&user)
			if user.ID == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized: ID not exist"})
				return
			}

			if user.Token != tokenString {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized: token string not match"})
				return
			}
		} else {
			log.Info("invalid token", zap.String("token", tokenString))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized: invalid token"})
			return
		}

		c.Next()
	}
}
