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

func (a *Appointment) GetByTimeSlot(slot internal.TimeSlot, employeeID int) ([]internal.Appointment, error) {
	sess := a.connection.NewSession(nil)

	var appointments []internal.Appointment
	_, err := sess.Select("*").
		From(appointmentsTable).
		Where(dbr.And(
			dbr.Eq("employee_id", employeeID),
			dbr.Or(
				// Appointment starts within the slot
				dbr.And(
					dbr.Gt("start_time", slot.StartTime),
					dbr.Lt("start_time", slot.EndTime),
				),
				// Appointment ends within the slot
				dbr.And(
					dbr.Gt("end_time", slot.StartTime),
					dbr.Lt("end_time", slot.EndTime),
				),
				// Appointment starts before and ends after the slot (full overlap)
				dbr.And(
					dbr.Lte("start_time", slot.StartTime),
					dbr.Gte("end_time", slot.EndTime),
				),
			),
		)).
		Load(&appointments)

	if err != nil {
		return nil, err
	}

	return appointments, nil
}
