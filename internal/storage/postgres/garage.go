package postgres

import (
	"strings"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/gocraft/dbr/v2"
)

const (
	garagesTable = "garages"
	pageSize     = 20
)

type Garage struct {
	connection *dbr.Connection
}

func NewGarage(connection *dbr.Connection) *Garage {
	return &Garage{
		connection: connection,
	}
}

func (g *Garage) Insert(garage internal.Garage) (internal.Garage, error) {
	sess := g.connection.NewSession(nil)

	var id int
	err := sess.InsertInto(garagesTable).
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

func (g *Garage) GetByOwnerID(employeeID int) (internal.Garage, error) {
	sess := g.connection.NewSession(nil)

	var garage internal.Garage
	err := sess.Select("*").
		From(garagesTable).
		Where(dbr.Eq("owner_id", employeeID)).
		LoadOne(&garage)

	return garage, err
}

func (g *Garage) GetByID(ID int) (internal.Garage, error) {
	sess := g.connection.NewSession(nil)

	var garage internal.Garage
	err := sess.Select("*").
		From(garagesTable).
		Where(dbr.Eq("id", ID)).
		LoadOne(&garage)

	return garage, err
}

func (g *Garage) List(query string, page int) ([]internal.Garage, error) {
	sess := g.connection.NewSession(nil)
	likeQuery := "%" + strings.ToLower(query) + "%"
	var garages []internal.Garage

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	_, err := sess.SelectBySql(`
        SELECT DISTINCT g.*
        FROM garages AS g
        LEFT JOIN services AS s ON s.garage_id = g.id
        WHERE LOWER(g.name) LIKE ? OR LOWER(s.name) LIKE ?
        LIMIT ?
        OFFSET ?
        `, likeQuery, likeQuery, pageSize, offset).
		Load(&garages)

	if err != nil {
		return []internal.Garage{}, err
	}

	return garages, nil
}
