package domain

import "time"

type Customer struct {
	ID        int64
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindByID(id int64) (*Customer, error)
	Create(customer *Customer) (*Customer, error)
	Update(customer *Customer) (*Customer, error)
	Delete(id int64) error
}
