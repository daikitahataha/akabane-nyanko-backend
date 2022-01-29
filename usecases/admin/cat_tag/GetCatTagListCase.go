package cat_tag

import (
	cts "akabane_nyanko/backend/src/services/admin/cat_tag_service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCatTagListCase struct {
	cts *cts.CatTagService
}

type CatTags = cts.CatTag

func New() *GetCatTagListCase {
	catTagService := cts.New()
	return &GetCatTagListCase{
		cts: catTagService,
	}
}

func (usecase GetCatTagListCase) Invoke(c *gin.Context) ([]CatTags, error) {
	page := c.Query("page")
	var p int
	p, _ = strconv.Atoi(page)

	catTags, err := usecase.cts.GetByPaginate(p)
	if err != nil {
		log.Fatalln(err)
	}

	return catTags, err
}
