package postgres

import (
	"github.com/KsaweryZietara/garage/internal"

	"github.com/gocraft/dbr/v2"
)

const servicesTable = "services"

type Service struct {
	connection *dbr.Connection
}

func NewService(connection *dbr.Connection) *Service {
	return &Service{
		connection: connection,
	}
}

func (e *Service) Insert(service internal.Service) (internal.Service, error) {
	session := e.connection.NewSession(nil)

	var id int
	err := session.InsertInto(servicesTable).
		Columns("name", "time", "price", "garage_id").
		Record(service).
		Returning("id").
		Load(&id)

	if err != nil {
		return internal.Service{}, err
	}

	service.ID = id
	return service, nil
}
