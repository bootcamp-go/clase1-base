package internal

import "github.com/bootcamp-go/clase1-base/internal/models"

type Service interface {
	GetOne(id int) (models.Product, error)
	Store(product models.Product) (models.Product, error)
	Update(id int, product models.Product) (models.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetOne(id int) (models.Product, error) {
	product, err := s.repository.GetOne(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *service) Store(product models.Product) (models.Product, error) {
	prod, err := s.repository.Store(product)

	if err != nil {
		return models.Product{}, err
	}

	return prod, nil
}

func (s *service) Update(id int, product models.Product) (models.Product, error) {
	return s.repository.Update(id, product)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
