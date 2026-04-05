package usecase

import "github.com/tasukuwatanabe/crm-app/internal/domain"

type CustomerUsecase struct {
	repo domain.CustomerRepository
}

func NewCustomerUsecase(repo domain.CustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{repo: repo}
}

func (u *CustomerUsecase) GetAll() ([]domain.Customer, error) {
	return u.repo.FindAll()
}

func (u *CustomerUsecase) GetByID(id int64) (*domain.Customer, error) {
	return u.repo.FindByID(id)
}

func (u *CustomerUsecase) Create(customer *domain.Customer) (*domain.Customer, error) {
	return u.repo.Create(customer)
}

func (u *CustomerUsecase) Update(customer *domain.Customer) (*domain.Customer, error) {
	return u.repo.Update(customer)
}

func (u *CustomerUsecase) Delete(id int64) error {
	return u.repo.Delete(id)
}
