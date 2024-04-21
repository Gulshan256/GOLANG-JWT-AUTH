package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Gulshan256/GOLANG-JWT-AUTH/initializers"
	"github.com/Gulshan256/GOLANG-JWT-AUTH/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	// get the name, email and password off req body
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the name, email and password",
		})

		return

	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})

		return
	}

	// create the user
	user := models.AuthUser{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the user",
		})

		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {

	// get the email and pass to the body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the name, email and password",
		})

		return

	}

	// look up request user
	var user models.AuthUser
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": " invalid email or password",
		})
		return
	}

	// comapre sent in pass with saved user pass hash

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": " invalid email or password",
		})
		return
	}
	// genrate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})

		return
	}

	// send it back to
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Validate(c *gin.Context) {

	user, _  := c.Get("user")


	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})

}
