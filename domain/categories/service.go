package categories

import "context"

type RepostiryInterface interface {
	GetAll(ctx context.Context) (categoryList []Category, err error)
}

type service struct {
	repository RepostiryInterface
}

func newService(repository RepostiryInterface) service {
	return service{
		repository: repository,
	}
}

func (s service) GetAll(ctx context.Context) (categoryList []Category, err error) {
	categoryList, err = s.repository.GetAll(ctx)
	if err != nil {
		if err == ErrCategoryNotFound {
			return []Category{}, err
		}
		return nil, err
	}

	return categoryList, nil
}
