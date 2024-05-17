package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/gin-auth/internal/models"
)

func BlogsIndex(c *gin.Context) {
	blogs := models.BlogsAll()

	user, _ := c.Get("user")

	c.HTML(
		http.StatusOK,
		"blogs/index.tpl",
		gin.H{
			"blogs": blogs,
			"user":  user.(models.User),
		},
	)
}

func BlgosShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	blog := models.BlogsFind(id)

	c.HTML(
		http.StatusOK,
		"blogs/show.tpl",
		gin.H{
			"blog": blog,
		},
	)
}
