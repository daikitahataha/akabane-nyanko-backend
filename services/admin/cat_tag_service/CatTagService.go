package cat_tag_service

import (
	"akabane_nyanko/backend/src/repositories/admin/cat_tag_repository"
	"log"
	"math"
)

type CatTagService struct {
	ctr cat_tag_repository.CatTagRepositoryInterface
}

type CatTag = cat_tag_repository.CatTag

func New(ctr cat_tag_repository.CatTagRepositoryInterface) *CatTagService {
	return &CatTagService{ctr}
}

func (s CatTagService) GetAsPaginate(page int) ([]CatTag, int, error) {
	var perPage = 10

	offset := (page - 1) * perPage
	limit := offset * perPage
	if offset == 0 {
		limit = perPage
	}

	catTags, err := s.ctr.GetByLimitAndOffset(offset, limit)
	if err != nil {
		log.Fatalln(err)
	}

	totalRecord := s.ctr.GetRecordCount()
	totalPageBefore := float64(totalRecord) / float64(perPage)
	totalPage := math.Ceil(totalPageBefore)

	return catTags, int(totalPage), err
}
