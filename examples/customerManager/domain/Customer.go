package domain

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerRepository interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	FindById(string) (Customer, error)
	FindAll() ([]Customer, error)
}
