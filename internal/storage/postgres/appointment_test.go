package postgres

import (
	"testing"
	"time"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestAppointment(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)
	serviceRepo := NewService(connection)
	customerRepo := NewCustomer(connection)
	appointmentRepo := NewAppointment(connection)

	newEmployee := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "test@test.com",
		Password: "password123",
		Role:     internal.OwnerRole,
	}
	employee, err := employeeRepo.Insert(newEmployee)
	assert.NoError(t, err)

	newGarage := internal.Garage{
		Name:        "Test Garage",
		City:        "Test City",
		Street:      "Test Street",
		Number:      "123",
		PostalCode:  "12345",
		PhoneNumber: "1234567890",
		OwnerID:     employee.ID,
	}
	garage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	newEmployee2 := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "test2@test.com",
		Password: "password123",
		Role:     internal.MechanicRole,
		GarageID: &garage.ID,
	}
	employee2, err := employeeRepo.Insert(newEmployee2)
	assert.NoError(t, err)

	newService := internal.Service{
		Name:     "Test Service",
		Time:     60,
		Price:    100.0,
		GarageID: garage.ID,
	}
	service, err := serviceRepo.Insert(newService)
	assert.NoError(t, err)

	newCustomer := internal.Customer{
		Email:    "test@test.com",
		Password: "password123",
	}
	customer, err := customerRepo.Insert(newCustomer)
	assert.NoError(t, err)

	startTime := time.Now()

	appointments := []internal.Appointment{
		{
			StartTime:  startTime.Add(1 * time.Hour),
			EndTime:    startTime.Add(2 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee2.ID,
			CustomerID: customer.ID,
		},
		{
			StartTime:  startTime.Add(2 * time.Hour),
			EndTime:    startTime.Add(4 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee2.ID,
			CustomerID: customer.ID,
		},
		{
			StartTime:  startTime.Add(4 * time.Hour),
			EndTime:    startTime.Add(5 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee2.ID,
			CustomerID: customer.ID,
		},
	}

	for _, appointment := range appointments {
		_, err = appointmentRepo.Insert(appointment)
		assert.NoError(t, err)
	}

	timeSlot := internal.TimeSlot{
		StartTime: startTime.Add(1 * time.Hour),
		EndTime:   startTime.Add(3 * time.Hour),
	}

	foundAppointments, err := appointmentRepo.GetByTimeSlot(timeSlot, employee2.ID)
	assert.NoError(t, err)

	assert.Len(t, foundAppointments, 2)

	assert.Equal(t, appointments[0].StartTime.Hour(), foundAppointments[0].StartTime.Hour())
	assert.Equal(t, appointments[1].StartTime.Hour(), foundAppointments[1].StartTime.Hour())

	nonOverlappingTimeSlot := internal.TimeSlot{
		StartTime: startTime.Add(6 * time.Hour),
		EndTime:   startTime.Add(7 * time.Hour),
	}

	foundAppointments, err = appointmentRepo.GetByTimeSlot(nonOverlappingTimeSlot, employee2.ID)
	assert.NoError(t, err)
	assert.Len(t, foundAppointments, 0)
}
