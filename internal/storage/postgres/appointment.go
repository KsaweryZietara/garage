package postgres

import (
	"time"

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

func (a *Appointment) GetByEmployeeID(employeeID int, date time.Time) ([]internal.Appointment, error) {
	sess := a.connection.NewSession(nil)

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	var appointments []internal.Appointment
	_, err := sess.Select("*").
		From(appointmentsTable).
		Where(dbr.And(
			dbr.Eq("employee_id", employeeID),
			dbr.Lte("start_time", endOfDay),
			dbr.Gte("end_time", startOfDay),
		)).
		OrderBy("start_time ASC").
		Load(&appointments)

	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (a *Appointment) GetByGarageID(garageID int, date time.Time) ([]internal.Appointment, error) {
	sess := a.connection.NewSession(nil)

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	var appointments []internal.Appointment
	_, err := sess.Select("a.*").
		From(dbr.I(appointmentsTable).As("a")).
		LeftJoin(dbr.I(employeesTable).As("e"), "a.employee_id = e.id").
		LeftJoin(dbr.I(garagesTable).As("g"), "e.garage_id = g.id").
		Where(dbr.And(
			dbr.Eq("g.id", garageID),
			dbr.Lte("start_time", endOfDay),
			dbr.Gte("end_time", startOfDay),
		)).
		OrderBy("a.start_time ASC").
		Load(&appointments)

	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (a *Appointment) GetByCustomerID(customerID int) ([]internal.Appointment, error) {
	sess := a.connection.NewSession(nil)

	var appointments []internal.Appointment
	_, err := sess.Select("*").
		From(appointmentsTable).
		Where(dbr.Eq("customer_id", customerID)).
		OrderBy("start_time DESC").
		Load(&appointments)

	if err != nil {
		return nil, err
	}

	return appointments, nil
}
