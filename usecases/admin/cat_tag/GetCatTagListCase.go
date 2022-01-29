package cat_tag

import (
	"akabane_nyanko/backend/src/repositories/admin/cat_tag_repository"
	"log"
)

type GetCatTagListCase struct {
	ctr cat_tag_repository.CatTagRepositoryInterface
}

type CatTags = cat_tag_repository.CatTag

func New() *GetCatTagListCase {
	catTagRepositoryBind := cat_tag_repository.GetBind()
	return &GetCatTagListCase{
		ctr: catTagRepositoryBind.Ctr,
	}
}

func (usecase GetCatTagListCase) Invoke(offset int, page int) ([]CatTags, error) {
	limit := page * 10
	catTags, err := usecase.ctr.GetByLimitAndOffset(offset, limit)
	if err != nil {
		log.Fatalln(err)
	}

	return catTags, err
}
