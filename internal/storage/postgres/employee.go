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
		Columns("name", "surname", "email", "password", "role", "garage_id", "confirmed").
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
		Where(dbr.And(
			dbr.Eq("email", email),
			dbr.Eq("confirmed", true),
		)).
		LoadOne(&employee)
	if err != nil {
		return internal.Employee{}, err
	}
	return employee, nil
}

func (e *Employee) Update(employee internal.Employee) error {
	sess := e.connection.NewSession(nil)
	_, err := sess.Update(employeesTable).
		Where(dbr.Eq("id", employee.ID)).
		Set("name", employee.Name).
		Set("surname", employee.Surname).
		Set("password", employee.Password).
		Set("confirmed", employee.Confirmed).
		Exec()

	return err
}

func (e *Employee) ListConfirmedByGarageID(garageID int) ([]internal.Employee, error) {
	sess := e.connection.NewSession(nil)

	var employees []internal.Employee
	_, err := sess.Select("*").
		From(employeesTable).
		Where(dbr.And(
			dbr.Eq("garage_id", garageID),
			dbr.Eq("confirmed", true),
		)).
		Load(&employees)

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *Employee) ListByGarageID(garageID int) ([]internal.Employee, error) {
	sess := e.connection.NewSession(nil)

	var employees []internal.Employee
	_, err := sess.Select("*").
		From(employeesTable).
		Where(dbr.Eq("garage_id", garageID)).
		Load(&employees)

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *Employee) GetConfirmedByID(ID int) (internal.Employee, error) {
	sess := e.connection.NewSession(nil)

	var employee internal.Employee
	err := sess.Select("*").
		From(employeesTable).
		Where(dbr.And(
			dbr.Eq("id", ID),
			dbr.Eq("confirmed", true),
		)).
		LoadOne(&employee)

	if err != nil {
		return internal.Employee{}, err
	}

	return employee, nil
}

func (e *Employee) GetByID(ID int) (internal.Employee, error) {
	sess := e.connection.NewSession(nil)

	var employee internal.Employee
	err := sess.Select("*").
		From(employeesTable).
		Where(dbr.Eq("id", ID)).
		LoadOne(&employee)

	if err != nil {
		return internal.Employee{}, err
	}

	return employee, nil
}
