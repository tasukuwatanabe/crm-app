package repository

import (
	"context"
	"database/sql"

	"github.com/tasukuwatanabe/crm-app/internal/db"
	"github.com/tasukuwatanabe/crm-app/internal/domain"
)

type postgresCustomerRepository struct {
	queries *db.Queries
}

func NewPostgresCustomerRepository(conn *sql.DB) domain.CustomerRepository {
	return &postgresCustomerRepository{
		queries: db.New(conn),
	}
}

func (r *postgresCustomerRepository) FindAll() ([]domain.Customer, error) {
	customers, err := r.queries.GetCustomers(context.Background())
	if err != nil {
		return nil, err
	}

	result := make([]domain.Customer, len(customers))
	for i, c := range customers {
		result[i] = toDomainCustomer(c)
	}

	return result, nil
}

func (r *postgresCustomerRepository) FindByID(id int64) (*domain.Customer, error) {
	c, err := r.queries.GetCustomersByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	customer := toDomainCustomer(c)
	return &customer, nil
}

func (r *postgresCustomerRepository) Create(customer *domain.Customer) (*domain.Customer, error) {
	c, err := r.queries.CreateCustomer(context.Background(), db.CreateCustomerParams{
		Name:  customer.Name,
		Email: customer.Email,
		Phone: sql.NullString{String: customer.Phone, Valid: customer.Phone != ""},
	})
	if err != nil {
		return nil, err
	}
	created := toDomainCustomer(c)
	return &created, nil
}

func (r *postgresCustomerRepository) Update(customer *domain.Customer) (*domain.Customer, error) {
	c, err := r.queries.UpdateCustomer(context.Background(), db.UpdateCustomerParams{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
		Phone: sql.NullString{String: customer.Phone, Valid: customer.Phone != ""},
	})
	if err != nil {
		return nil, err
	}
	updated := toDomainCustomer(c)
	return &updated, nil
}

func (r *postgresCustomerRepository) Delete(id int64) error {
	return r.queries.DeleteCustomer(context.Background(), id)
}

func toDomainCustomer(c db.Customer) domain.Customer {
	return domain.Customer{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone.String,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
