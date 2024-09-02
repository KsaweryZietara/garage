package postgres

import (
	"github.com/KsaweryZietara/garage/internal"
	"github.com/gocraft/dbr/v2"
)

const employeesTable = "employees"

type Employee struct {
	connection *dbr.Connection
}

func NewEmployee(connection *dbr.Connection) *Employee {
	return &Employee{
		connection: connection,
	}
}

func (e *Employee) Insert(employee internal.Employee) (internal.Employee, error) {
	sess := e.connection.NewSession(nil)
	var id int
	err := sess.InsertInto(employeesTable).
		Columns("name", "surname", "email", "password", "role", "garage_id").
		Record(employee).
		Returning("id").
		Load(&id)

	if err != nil {
		return internal.Employee{}, err
	}

	employee.ID = id
	return employee, nil
}

func (e *Employee) GetByEmail(email string) (internal.Employee, error) {
	var employee internal.Employee
	sess := e.connection.NewSession(nil)
	err := sess.Select("*").
		From(employeesTable).
		Where(dbr.Eq("email", email)).
		LoadOne(&employee)
	if err != nil {
		return internal.Employee{}, err
	}
	return employee, nil
}
