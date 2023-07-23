package server

import (
	"bebrah/app/db"
	"bebrah/app/middleware"
	"bebrah/app/model"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// @Description register
// @Tags auth
// @Param request body model.RegisterReq true "register request"
// @Success 200 {string} success
// @Router /auth/register [post]
func register(c *gin.Context) {
	var req model.RegisterReq
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

	db.Db().Create(&model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	})

	c.String(http.StatusOK, "success")
}

// @BasePath /api/v1

// @Description login
// @Tags auth
// @Param request body model.LoginReq true "login request"
// @Success 200 {object} model.LoginResp
// @Router /auth/login [post]
func login(c *gin.Context) {
	var req model.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var user model.User
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
		Id:        fmt.Sprintf("%d", user.ID),
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString(middleware.SecretKey)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	user.Token = tokenString
	db.Db().Model(&user).Update("token", tokenString)

	c.JSON(http.StatusOK, &model.LoginResp{
		Token:          tokenString,
		TokenExpiredAt: &expiredAt,
	})
}

// @BasePath /api/v1

// @Description logout
// @Tags auth
// @Success 200 {string} success
// @Router /auth/logout [post]
func logout(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var user model.User
	db.Db().Where("email = ?", userId).First(&user)
	user.Token = ""
	db.Db().Model(&user).Update("token", "")
	c.String(http.StatusOK, "success")
}

func setupAuth(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/register", register)
	auth.POST("/login", login)
	auth.POST("/logout", middleware.JWTAuthMiddleware(), logout)
	auth.POST("/verifyemail/:verifycationCode", func(c *gin.Context) {
		// TODO: implement
	})
}
