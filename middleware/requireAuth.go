package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Gulshan256/GOLANG-JWT-AUTH/initializers"
	"github.com/Gulshan256/GOLANG-JWT-AUTH/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	// get the cookie off request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {

		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECERET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {

		// ch3eck for exp
		if float64(time.Now().Unix()) > claims["exp"].(float64){

			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token 
		var user models.AuthUser
		initializers.DB.First(&user,  claims["sub"])

		if user .Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to request
		c.Set("user", user)


		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	}

	
}
