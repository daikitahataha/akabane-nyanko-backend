package cat_tag

import (
	"akabane_nyanko/backend/src/db"
	"akabane_nyanko/backend/src/models"
)

type CatTagRepository struct{}

type CatTag models.CatTag

func (ctr CatTagRepository) GetByLimitAndOffset(offset int, limit int) ([]CatTag, error) {
	db := db.GetDB()
	var catTags []CatTag

	if err := db.Select([]string{"ID", "Name"}).Limit(limit).Offset(offset).Find(&catTags).Error; err != nil {
		return nil, err
	}

	return catTags, nil
}
