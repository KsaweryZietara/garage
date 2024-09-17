package postgres

import (
	"github.com/KsaweryZietara/garage/internal"

	"github.com/gocraft/dbr/v2"
)

const customersTable = "customers"

type Customer struct {
	connection *dbr.Connection
}

func NewCustomer(connection *dbr.Connection) *Customer {
	return &Customer{
		connection: connection,
	}
}

func (c *Customer) Insert(customer internal.Customer) (internal.Customer, error) {
	sess := c.connection.NewSession(nil)
	var id int
	err := sess.InsertInto(customersTable).
		Columns("email", "password").
		Record(customer).
		Returning("id").
		Load(&id)

	if err != nil {
		return internal.Customer{}, err
	}

	customer.ID = id
	return customer, nil
}

func (c *Customer) GetByEmail(email string) (internal.Customer, error) {
	var customer internal.Customer
	sess := c.connection.NewSession(nil)
	err := sess.Select("*").
		From(customersTable).
		Where(dbr.Eq("email", email)).
		LoadOne(&customer)

	if err != nil {
		return internal.Customer{}, err
	}

	return customer, nil
}
