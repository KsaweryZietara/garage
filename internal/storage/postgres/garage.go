package postgres

import (
	"github.com/KsaweryZietara/garage/internal"

	"github.com/gocraft/dbr/v2"
)

const garagesTable = "garages"

type Garage struct {
	connection *dbr.Connection
}

func NewGarage(connection *dbr.Connection) *Garage {
	return &Garage{
		connection: connection,
	}
}

func (e *Garage) Insert(garage internal.Garage) (internal.Garage, error) {
	session := e.connection.NewSession(nil)

	var id int
	err := session.InsertInto(garagesTable).
		Columns("name", "city", "street", "number", "postal_code", "phone_number", "owner_id").
		Record(garage).
		Returning("id").
		Load(&id)

	if err != nil {
		return internal.Garage{}, err
	}

	garage.ID = id
	return garage, nil
}

func (e *Garage) GetByOwnerID(employeeID int) (internal.Garage, error) {
	session := e.connection.NewSession(nil)

	var garage internal.Garage
	err := session.Select("*").
		From(garagesTable).
		Where(dbr.Eq("owner_id", employeeID)).
		LoadOne(&garage)

	return garage, err
}

func (e *Garage) GetByID(ID int) (internal.Garage, error) {
	session := e.connection.NewSession(nil)

	var garage internal.Garage
	err := session.Select("*").
		From(garagesTable).
		Where(dbr.Eq("id", ID)).
		LoadOne(&garage)

	return garage, err
}
