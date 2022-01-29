package admin

import (
	"net/http"

	CatTagCase "akabane_nyanko/backend/src/usecases/admin/cat_tag"

	"github.com/gin-gonic/gin"
)

type CatTagController struct{}

func (ctc CatTagController) TagList(c *gin.Context) {
	usecase := CatTagCase.New()
	catTags, err := usecase.Invoke(c)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{
			"message": "ng",
			"data":    catTags,
			"error":   err,
		})
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    catTags,
	})
}
