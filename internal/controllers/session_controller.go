package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leetcode-golang-classroom/gin-auth/internal/config"
	"github.com/leetcode-golang-classroom/gin-auth/internal/models"
)

func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"sessions/signup.tpl",
		gin.H{},
	)
}

func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"sessions/login.tpl",
		gin.H{},
	)
}

type formData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Signup(c *gin.Context) {
	var data formData
	c.Bind(&data)

	// check if the user exists already
	if !models.CheckUserAvailability(data.Email) {
		c.HTML(http.StatusBadGateway,
			"errors/error.tpl",
			gin.H{"error": "Email missing"},
		)
		return
	}
	// Create the user
	user := models.UserCreate(data.Email, data.Password)
	if user == nil || user.ID == 0 {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.tpl",
			gin.H{"error": "user creation failed"},
		)
		return
	}
	// create a jwt
	tokenString, err := createAndSignJWT(user)
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.tpl",
			gin.H{"error": "JWT creation failed"},
		)
		return
	}

	// 2 send the token in a cookie
	setCookie(c, tokenString)
	c.Redirect(http.StatusFound, "/blogs")
}

func Login(c *gin.Context) {
	var data formData
	c.Bind(&data)
	// Match password
	user := models.UserMatchPassword(data.Email, data.Password)
	if user.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"errors/error.tpl",
			gin.H{"error": "Unauthorized"},
		)
		return
	}
	// Create JWT token
	tokenString, err := createAndSignJWT(user)
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.tpl",
			gin.H{"error": "JWT creation failed"},
		)
		return
	}
	// Send the token in cookie
	setCookie(c, tokenString)
	c.Redirect(http.StatusFound, "/blogs")
}
func Logout(c *gin.Context) {
	// Add the JWT token to the block list
	// or change expiry time of  the cookie

	c.SetCookie("Auth", "deleted", 0, "", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}

func createAndSignJWT(user *models.User) (string, error) {
	// 1 create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"ttl":    time.Now().Add(time.Hour * 24 * 100).Unix(),
	})
	// load secret from environment
	hmacSampleSecret := config.C.JWT_SIGN_SECRET
	// sign and get complete signed token
	return token.SignedString([]byte(hmacSampleSecret))
}

func setCookie(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", token, 3600*24*100, "", "", false, true)
}
