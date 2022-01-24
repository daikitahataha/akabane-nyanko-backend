package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TagList(c *gin.Context) {
	fmt.Println("hello")
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    "tag",
	})
}
