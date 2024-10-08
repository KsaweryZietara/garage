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
		Columns("name", "city", "street", "number", "postal_code", "phone_number", "latitude", "longitude", "owner_id").
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

func (g *Garage) List(page int, query string, latitude, longitude float64) ([]internal.Garage, error) {
	sess := g.connection.NewSession(nil)
	likeQuery := "%" + strings.ToLower(query) + "%"
	var garages []internal.Garage
	var err error

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	if latitude != 0 && longitude != 0 {
		_, err = sess.SelectBySql(`
        SELECT DISTINCT g.*, COALESCE(AVG(a.rating), 0) AS rating,
            COALESCE(( 6371 * acos( cos( radians(?) ) * cos( radians(g.latitude) ) 
            * cos( radians(g.longitude) - radians(?) ) + sin( radians(?) ) 
            * sin( radians(g.latitude) ) ) ), 0) AS distance
        FROM garages AS g
        LEFT JOIN services AS s ON s.garage_id = g.id
        LEFT JOIN employees AS e ON e.garage_id = g.id
        LEFT JOIN appointments AS a ON a.employee_id = e.id
        WHERE LOWER(g.name) LIKE ? OR LOWER(s.name) LIKE ?
        GROUP BY g.id, g.name, city, street, number, postal_code, phone_number, owner_id
		ORDER BY distance
        LIMIT ?
        OFFSET ?
        `, latitude, longitude, latitude, likeQuery, likeQuery, pageSize, offset).
			Load(&garages)
	} else {
		_, err = sess.SelectBySql(`
        SELECT DISTINCT g.*, COALESCE(AVG(a.rating), 0) AS rating
        FROM garages AS g
        LEFT JOIN services AS s ON s.garage_id = g.id
        LEFT JOIN employees AS e ON e.garage_id = g.id
        LEFT JOIN appointments AS a ON a.employee_id = e.id
        WHERE LOWER(g.name) LIKE ? OR LOWER(s.name) LIKE ?
        GROUP BY g.id, g.name, city, street, number, postal_code, phone_number, owner_id
		ORDER BY rating DESC
        LIMIT ?
        OFFSET ?
        `, likeQuery, likeQuery, pageSize, offset).
			Load(&garages)
	}

	if err != nil {
		return []internal.Garage{}, err
	}

	return garages, nil
}
