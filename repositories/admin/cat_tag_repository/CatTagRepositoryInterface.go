package cat_tag_repository

type CatTagRepositoryInterface interface {
	GetByLimitAndOffset(offset int, limit int) ([]CatTag, error)
}

type Bind struct {
	Ctr CatTagRepositoryInterface
}

func NewBind() *Bind {
	var Ctr CatTagRepositoryInterface = &CatTagRepository{}
	return &Bind{
		Ctr,
	}
}

func GetBind() *Bind {
	return NewBind()
}
