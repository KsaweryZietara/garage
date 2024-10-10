package postgres

import (
	"github.com/KsaweryZietara/garage/internal"

	"github.com/gocraft/dbr/v2"
)

const (
	makesTable  = "makes"
	modelsTable = "models"
)

type Car struct {
	connection *dbr.Connection
}

func NewCar(connection *dbr.Connection) *Car {
	return &Car{
		connection: connection,
	}
}

func (c *Car) ListMakes() ([]internal.Make, error) {
	sess := c.connection.NewSession(nil)

	var makes []internal.Make
	_, err := sess.Select("*").
		From(makesTable).
		OrderBy("name").
		Load(&makes)

	if err != nil {
		return nil, err
	}

	return makes, nil
}

func (c *Car) ListModels(makeID int) ([]internal.Model, error) {
	sess := c.connection.NewSession(nil)

	var models []internal.Model
	_, err := sess.Select("*").
		From(modelsTable).
		Where(dbr.Eq("make_id", makeID)).
		OrderBy("name").
		Load(&models)

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (c *Car) GetByModelID(modelID int) (internal.Car, error) {
	sess := c.connection.NewSession(nil)

	var car internal.Car
	_, err := sess.Select("ma.name AS make, mo.name AS model").
		From(dbr.I(makesTable).As("ma")).
		Join(dbr.I(modelsTable).As("mo"), "mo.make_id = ma.id").
		Where(dbr.Eq("mo.id", modelID)).
		Load(&car)

	if err != nil {
		return internal.Car{}, err
	}

	return car, nil
}
