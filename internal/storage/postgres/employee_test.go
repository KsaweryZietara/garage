package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestEmployee(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)

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

	retrievedEmployee2, err := employeeRepo.GetConfirmedByID(employee2.ID)
	assert.NoError(t, err)
	assert.Equal(t, newEmployee2.Name, retrievedEmployee2.Name)
	assert.Equal(t, newEmployee2.Surname, retrievedEmployee2.Surname)
	assert.Equal(t, newEmployee2.Email, retrievedEmployee2.Email)
	assert.Equal(t, newEmployee2.Password, retrievedEmployee2.Password)
	assert.Equal(t, newEmployee2.Role, retrievedEmployee2.Role)

	retrievedEmployee2, err = employeeRepo.GetByID(employee2.ID)
	assert.NoError(t, err)
	assert.Equal(t, newEmployee2.Name, retrievedEmployee2.Name)
	assert.Equal(t, newEmployee2.Surname, retrievedEmployee2.Surname)
	assert.Equal(t, newEmployee2.Email, retrievedEmployee2.Email)
	assert.Equal(t, newEmployee2.Password, retrievedEmployee2.Password)
	assert.Equal(t, newEmployee2.Role, retrievedEmployee2.Role)

	newEmployee3 := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test3@test.com",
		Password:  "password123",
		Role:      internal.MechanicRole,
		GarageID:  &garage.ID,
		Confirmed: true,
	}
	_, err = employeeRepo.Insert(newEmployee3)
	assert.NoError(t, err)

	retrievedEmployee, err := employeeRepo.GetByEmail(newEmployee.Email)
	assert.NoError(t, err)
	assert.Equal(t, newEmployee.Name, retrievedEmployee.Name)
	assert.Equal(t, newEmployee.Surname, retrievedEmployee.Surname)
	assert.Equal(t, newEmployee.Email, retrievedEmployee.Email)
	assert.Equal(t, newEmployee.Password, retrievedEmployee.Password)
	assert.Equal(t, newEmployee.Role, retrievedEmployee.Role)

	employee.Name = "newName"
	employee.Surname = "newSurname"
	employee.Password = "newPassword"
	err = employeeRepo.Update(employee)
	assert.NoError(t, err)

	updatedEmployee, err := employeeRepo.GetByEmail(newEmployee.Email)
	assert.NoError(t, err)
	assert.Equal(t, employee.Name, updatedEmployee.Name)
	assert.Equal(t, employee.Surname, updatedEmployee.Surname)
	assert.Equal(t, employee.Password, updatedEmployee.Password)

	employees, err := employeeRepo.ListConfirmedByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(employees))

	newEmployee4 := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "test4@test.com",
		Password:  "password123",
		Role:      internal.MechanicRole,
		GarageID:  &garage.ID,
		Confirmed: false,
	}
	employee4, err := employeeRepo.Insert(newEmployee4)
	assert.NoError(t, err)

	_, err = employeeRepo.GetConfirmedByID(employee4.ID)
	assert.EqualError(t, err, "dbr: not found")

	_, err = employeeRepo.GetByEmail(employee4.Email)
	assert.EqualError(t, err, "dbr: not found")

	employees, err = employeeRepo.ListConfirmedByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(employees))

	employees, err = employeeRepo.ListByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(employees))

	employee4.Confirmed = true
	err = employeeRepo.Update(employee4)
	assert.NoError(t, err)

	employees, err = employeeRepo.ListConfirmedByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(employees))

	employees, err = employeeRepo.ListByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(employees))

	err = employeeRepo.Delete(employee2.ID)
	assert.NoError(t, err)

	deletedEmployee, err := employeeRepo.GetByID(employee2.ID)
	assert.NoError(t, err)
	assert.True(t, deletedEmployee.IsDeleted)

	_, err = employeeRepo.GetByEmail(employee2.Email)
	assert.EqualError(t, err, "dbr: not found")

	employees, err = employeeRepo.ListConfirmedByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(employees))

	employees, err = employeeRepo.ListByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(employees))

	profilePicture := []byte("profile-picture")
	err = employeeRepo.UpdateProfilePicture(employee4.ID, profilePicture)
	assert.NoError(t, err)

	employee, err = employeeRepo.GetByID(employee4.ID)
	assert.NoError(t, err)
	assert.Equal(t, profilePicture, employee.ProfilePicture)
}
