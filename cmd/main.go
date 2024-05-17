package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/gin-auth/internal/config"
	"github.com/leetcode-golang-classroom/gin-auth/internal/models"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("internal/templates/**/*")
	models.ConnectDatabase(config.C.MYSQL_URI)
	models.DBMigrate()
	log.Printf("Server started at port:%v", config.C.SERVER_PORT)
	r.Run(fmt.Sprintf(":%v", config.C.SERVER_PORT))
}
