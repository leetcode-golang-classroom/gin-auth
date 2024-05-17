package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/gin-auth/internal/config"
	"github.com/leetcode-golang-classroom/gin-auth/internal/controllers"
	"github.com/leetcode-golang-classroom/gin-auth/internal/middlewares"
	"github.com/leetcode-golang-classroom/gin-auth/internal/models"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	// setup template
	r.LoadHTMLGlob("internal/templates/**/*")
	// setup gorm
	models.ConnectDatabase(config.C.MYSQL_URI)
	models.DBMigrate()

	r.GET("/blogs", middlewares.AuthMiddleware, controllers.BlogsIndex)
	r.GET("/blogs/:id", middlewares.AuthMiddleware, controllers.BlogsIndex)

	r.GET("/signup", controllers.SignupPage)
	r.GET("/login", controllers.LoginPage)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.DELETE("/logout", controllers.Logout)

	log.Printf("Server started at port:%v", config.C.SERVER_PORT)
	r.Run(fmt.Sprintf(":%v", config.C.SERVER_PORT))
}
