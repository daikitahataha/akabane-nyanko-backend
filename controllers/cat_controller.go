package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CatList(c *gin.Context) {
	fmt.Println("hello")
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"catData": "nyaa",
	})
}
