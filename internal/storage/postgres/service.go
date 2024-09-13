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

func (s *Service) Insert(service internal.Service) (internal.Service, error) {
	sess := s.connection.NewSession(nil)

	var id int
	err := sess.InsertInto(servicesTable).
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

func (s *Service) ListByGarageID(garageID int) ([]internal.Service, error) {
	sess := s.connection.NewSession(nil)

	var services []internal.Service
	_, err := sess.Select("*").
		From(servicesTable).
		Where(dbr.Eq("garage_id", garageID)).
		Load(&services)

	if err != nil {
		return nil, err
	}

	return services, nil
}
