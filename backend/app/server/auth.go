package server

import (
	"bebrah/app/db"
	"bebrah/app/middleware"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token          string     `json:"token"`
	TokenExpiredAt *time.Time `json:"tokenExpiredAt"`
}

func setupAuth(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/register", func(c *gin.Context) {
		var req RegisterReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		if req.Password != req.ConfirmPassword {
			c.JSON(http.StatusBadRequest, gin.H{"message": "passwords do not match"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		}

		db.Db().Create(&db.User{
			Email:    req.Email,
			Password: string(hashedPassword),
		})

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	auth.POST("/login", func(c *gin.Context) {
		var req LoginReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		var user db.User
		db.Db().Where("email = ?", req.Email).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email or password"})
			return
		}

		expiredAt := time.Now().Add(time.Hour * 3)
		token := &jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			Issuer:    user.Email,
		}

		tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString(middleware.SecretKey)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			return
		}

		user.Token = tokenString
		db.Db().Model(&user).Update("token", tokenString)

		c.JSON(http.StatusOK, &LoginResp{
			Token:          tokenString,
			TokenExpiredAt: &expiredAt,
		})
	})
	auth.POST("/logout", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		userId, exist := c.Get("user")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request, user not found"})
			return
		}

		var user db.User
		db.Db().Where("id = ?", userId).First(&user)
		user.Token = ""
		db.Db().Model(&user).Update("token", "")
	})
	auth.POST("/verifyemail/:verifycationCode", func(c *gin.Context) {
		// TODO: implement
	})
}
