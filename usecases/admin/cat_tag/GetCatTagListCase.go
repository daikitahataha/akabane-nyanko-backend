package cat_tag

import (
	"akabane_nyanko/backend/src/repositories/admin/cat_tag"
	"log"
)

type GetCatTagListCase struct {
	ctr cat_tag.CatTagRepository
}

type CatTags = cat_tag.CatTag

func (usecase GetCatTagListCase) Invoke(offset int, page int) ([]CatTags, error) {
	limit := page * 10
	catTags, err := usecase.ctr.GetByLimitAndOffset(offset, limit)
	if err != nil {
		log.Fatalln(err)
	}

	return catTags, err
}
