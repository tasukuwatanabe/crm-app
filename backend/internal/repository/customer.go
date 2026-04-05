package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/tasukuwatanabe/crm-app/internal/domain"
)

type inMemoryCustomerRepository struct {
	mu       sync.Mutex
	store    map[int64]domain.Customer
	sequence int64
}

func NewInMemoryCustomerRepository() domain.CustomerRepository {
	return &inMemoryCustomerRepository{
		store: make(map[int64]domain.Customer),
	}
}

func (r *inMemoryCustomerRepository) FindAll() ([]domain.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	customers := make([]domain.Customer, 0, len(r.store))
	for _, c := range r.store {
		customers = append(customers, c)
	}
	return customers, nil
}

func (r *inMemoryCustomerRepository) FindByID(id int64) (*domain.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	c, ok := r.store[id]
	if !ok {
		return nil, errors.New("customer not found")
	}
	return &c, nil
}

func (r *inMemoryCustomerRepository) Create(customer *domain.Customer) (*domain.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.sequence++
	customer.ID = r.sequence
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	r.store[customer.ID] = *customer
	return customer, nil
}

func (r *inMemoryCustomerRepository) Update(customer *domain.Customer) (*domain.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.store[customer.ID]; !ok {
		return nil, errors.New("customer not found")
	}
	customer.UpdatedAt = time.Now()
	r.store[customer.ID] = *customer
	return customer, nil
}

func (r *inMemoryCustomerRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.store[id]; !ok {
		return errors.New("customer not found")
	}
	delete(r.store, id)
	return nil
}
