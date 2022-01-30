package cat_tag_service

import (
	"akabane_nyanko/backend/src/repositories/admin/cat_tag_repository"
	"log"
)

type CatTagService struct {
	ctr cat_tag_repository.CatTagRepositoryInterface
}

type CatTag = cat_tag_repository.CatTag

func New() *CatTagService {
	catTagRepositoryBind := cat_tag_repository.GetBind()
	return &CatTagService{
		ctr: catTagRepositoryBind.Ctr,
	}
}

func (s CatTagService) GetByPaginate(page interface{}) ([]CatTag, error) {
	p := page.(int)
	offset := p - 1
	limit := p * 10

	catTags, err := s.ctr.GetByLimitAndOffset(offset, limit)
	if err != nil {
		log.Fatalln(err)
	}

	return catTags, err
}
