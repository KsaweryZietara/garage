package postgres

import (
	"github.com/KsaweryZietara/garage/internal"

	"github.com/gocraft/dbr/v2"
)

const appointmentsTable = "appointments"

type Appointment struct {
	connection *dbr.Connection
}

func NewAppointment(connection *dbr.Connection) *Appointment {
	return &Appointment{
		connection: connection,
	}
}

func (a *Appointment) Insert(appointment internal.Appointment) (internal.Appointment, error) {
	sess := a.connection.NewSession(nil)

	var id int
	err := sess.InsertInto(appointmentsTable).
		Columns("start_time", "end_time", "service_id", "employee_id", "customer_id").
		Record(appointment).
		Returning("id").
		Load(&id)

	if err != nil {
		return internal.Appointment{}, err
	}

	appointment.ID = id
	return appointment, nil
}
