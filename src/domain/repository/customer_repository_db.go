package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/igson/banking/src/domain/models"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

//AccountRepository acesso ao repositório
type customerRepository struct {
	client *sqlx.DB
}

//NewCustomerRepository acesso ao repositório
func NewCustomerRepository(dbClient *sqlx.DB) interfaces.ICustomerRepository {
	return &customerRepository{
		client: dbClient,
	}
}

func (d customerRepository) FindAll(status string) ([]models.Customer, *errors.RestErroAPI) {
	var err error
	customers := make([]models.Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d customerRepository) ById(id int64) (*models.Customer, *errors.RestErroAPI) {

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c models.Customer

	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundErro("Customer not found")
		} else {
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}
