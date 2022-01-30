package cat_tag_service

import (
	"akabane_nyanko/backend/src/repositories/admin/cat_tag_repository"
	"log"
)

type CatTagService struct {
	ctr cat_tag_repository.CatTagRepositoryInterface
}

type CatTag = cat_tag_repository.CatTag

func New(ctr cat_tag_repository.CatTagRepositoryInterface) *CatTagService {
	return &CatTagService{ctr}
}

func (s CatTagService) GetByPaginate(page int) ([]CatTag, error) {
	offset := page - 1
	limit := page * 10

	catTags, err := s.ctr.GetByLimitAndOffset(offset, limit)
	if err != nil {
		log.Fatalln(err)
	}

	return catTags, err
}
