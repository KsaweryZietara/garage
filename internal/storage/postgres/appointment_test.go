package postgres

import (
	"testing"
	"time"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAppointmentsByTimeSlot(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)
	serviceRepo := NewService(connection)
	customerRepo := NewCustomer(connection)
	appointmentRepo := NewAppointment(connection)

	newEmployee := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test@test.com",
		Password:  "password123",
		Role:      internal.OwnerRole,
		Confirmed: true,
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
		Latitude:    10,
		Longitude:   10,
	}
	garage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	newEmployee2 := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test2@test.com",
		Password:  "password123",
		Role:      internal.MechanicRole,
		GarageID:  &garage.ID,
		Confirmed: true,
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
			ModelID:    1,
		},
		{
			StartTime:  startTime.Add(2 * time.Hour),
			EndTime:    startTime.Add(4 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee2.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		},
		{
			StartTime:  startTime.Add(4 * time.Hour),
			EndTime:    startTime.Add(5 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee2.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		},
	}

	for i, appointment := range appointments {
		insertedAppointment, err := appointmentRepo.Insert(appointment)
		assert.NoError(t, err)
		if i == 1 {
			rating := 4
			insertedAppointment.Rating = &rating
			err = appointmentRepo.Update(insertedAppointment)
			assert.NoError(t, err)
		}
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

	foundAppointments, err = appointmentRepo.ListByGarageID(garage.ID)
	require.NoError(t, err)
	assert.Len(t, foundAppointments, 1)
}

func TestGetAppointmentsByEmployeeIDOrCustomerID(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)
	serviceRepo := NewService(connection)
	customerRepo := NewCustomer(connection)
	appointmentRepo := NewAppointment(connection)

	newEmployee := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test@test.com",
		Password:  "password123",
		Role:      internal.OwnerRole,
		Confirmed: true,
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
		Latitude:    10,
		Longitude:   10,
	}
	garage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	newEmployee2 := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test2@test.com",
		Password:  "password123",
		Role:      internal.MechanicRole,
		GarageID:  &garage.ID,
		Confirmed: true,
	}
	employee2, err := employeeRepo.Insert(newEmployee2)
	assert.NoError(t, err)

	newEmployee3 := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test3@test.com",
		Password:  "password123",
		Role:      internal.MechanicRole,
		GarageID:  &garage.ID,
		Confirmed: true,
	}
	employee3, err := employeeRepo.Insert(newEmployee3)
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

	newCustomer2 := internal.Customer{
		Email:    "test2@test.com",
		Password: "password123",
	}
	customer2, err := customerRepo.Insert(newCustomer2)
	assert.NoError(t, err)

	appointments := []internal.Appointment{
		{
			StartTime:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()).Add(-50 * time.Hour),
			EndTime:    time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()).Add(-50 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee2.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		},
		{
			StartTime:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()).Add(-50 * time.Hour),
			EndTime:    time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()),
			ServiceID:  service.ID,
			EmployeeID: employee3.ID,
			CustomerID: customer2.ID,
			ModelID:    1,
		},
		{
			StartTime:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()),
			EndTime:    time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 0, 0, 0, time.Now().Location()),
			ServiceID:  service.ID,
			EmployeeID: employee2.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		},
		{
			StartTime:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()),
			EndTime:    time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()).Add(50 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee3.ID,
			CustomerID: customer2.ID,
			ModelID:    1,
		},
		{
			StartTime:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()).Add(50 * time.Hour),
			EndTime:    time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()).Add(50 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: employee3.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		},
	}

	for _, appointment := range appointments {
		insertedAppointment, err := appointmentRepo.Insert(appointment)
		assert.NoError(t, err)

		retrievedAppointment, err := appointmentRepo.GetByID(insertedAppointment.ID)
		assert.NoError(t, err)
		assert.Nil(t, retrievedAppointment.Rating)
		assert.Nil(t, retrievedAppointment.Comment)

		rating := 4
		comment := "comment"
		retrievedAppointment.Rating = &rating
		retrievedAppointment.Comment = &comment
		err = appointmentRepo.Update(retrievedAppointment)
		assert.NoError(t, err)

		updatedAppointment, err := appointmentRepo.GetByID(insertedAppointment.ID)
		assert.NoError(t, err)
		assert.Equal(t, retrievedAppointment.Rating, updatedAppointment.Rating)
		assert.Equal(t, retrievedAppointment.Comment, updatedAppointment.Comment)

		updatedAppointment.Rating = nil
		updatedAppointment.Comment = nil
		err = appointmentRepo.Update(updatedAppointment)
		assert.NoError(t, err)

		retrievedAppointment, err = appointmentRepo.GetByID(insertedAppointment.ID)
		assert.NoError(t, err)
		assert.Nil(t, retrievedAppointment.Rating)
		assert.Nil(t, retrievedAppointment.Comment)
	}

	foundAppointments, err := appointmentRepo.GetByEmployeeID(employee3.ID, time.Now())
	require.NoError(t, err)
	assert.Len(t, foundAppointments, 2)

	foundAppointments, err = appointmentRepo.GetByGarageID(garage.ID, time.Now())
	require.NoError(t, err)
	assert.Len(t, foundAppointments, 3)

	foundAppointments, err = appointmentRepo.GetByCustomerID(customer.ID)
	require.NoError(t, err)
	assert.Len(t, foundAppointments, 3)

	foundAppointments, err = appointmentRepo.GetByCustomerID(customer2.ID)
	require.NoError(t, err)
	assert.Len(t, foundAppointments, 2)
}

func TestDeleteAppointment(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)
	serviceRepo := NewService(connection)
	customerRepo := NewCustomer(connection)
	appointmentRepo := NewAppointment(connection)

	newEmployee := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test@test.com",
		Password:  "password123",
		Role:      internal.OwnerRole,
		Confirmed: true,
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
		Latitude:    10,
		Longitude:   10,
	}
	garage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	newEmployee2 := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test2@test.com",
		Password:  "password123",
		Role:      internal.MechanicRole,
		GarageID:  &garage.ID,
		Confirmed: true,
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

	newAppointment := internal.Appointment{
		StartTime:  time.Now(),
		EndTime:    time.Now(),
		ServiceID:  service.ID,
		EmployeeID: employee2.ID,
		CustomerID: customer.ID,
		ModelID:    1,
	}

	appointment, err := appointmentRepo.Insert(newAppointment)
	assert.NoError(t, err)

	_, err = appointmentRepo.GetByID(appointment.ID)
	assert.NoError(t, err)

	err = appointmentRepo.Delete(appointment.ID)
	assert.NoError(t, err)

	_, err = appointmentRepo.GetByID(appointment.ID)
	assert.EqualError(t, err, "dbr: not found")
}
