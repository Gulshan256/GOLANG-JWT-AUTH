package main

import (
	// "fmt"

	// "net/http"

	"fmt"

	"github.com/Gulshan256/GOLANG-JWT-AUTH/controllers"
	"github.com/Gulshan256/GOLANG-JWT-AUTH/initializers"
	"github.com/Gulshan256/GOLANG-JWT-AUTH/middleware"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvvariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/Validate", middleware.RequireAuth, controllers.Validate)

	fmt.Print("Server is running on port 8080")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
