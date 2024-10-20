package postgres

import (
	"github.com/KsaweryZietara/garage/internal"
	"github.com/gocraft/dbr/v2"
)

const confirmationCodesTable = "confirmation_codes"

type ConfirmationCode struct {
	connection *dbr.Connection
}

func NewConfirmationCode(connection *dbr.Connection) *ConfirmationCode {
	return &ConfirmationCode{
		connection: connection,
	}
}

func (c *ConfirmationCode) Insert(code internal.ConfirmationCode) (internal.ConfirmationCode, error) {
	sess := c.connection.NewSession(nil)
	_, err := sess.InsertInto(confirmationCodesTable).
		Columns("id", "employee_id").
		Record(code).
		Exec()

	if err != nil {
		return internal.ConfirmationCode{}, err
	}

	return code, nil
}

func (c *ConfirmationCode) GetByID(ID string) (internal.ConfirmationCode, error) {
	sess := c.connection.NewSession(nil)
	var code internal.ConfirmationCode
	err := sess.Select("*").
		From(confirmationCodesTable).
		Where(dbr.Eq("id", ID)).
		LoadOne(&code)

	return code, err
}

func (c *ConfirmationCode) DeleteByEmployeeID(ID int) error {
	sess := c.connection.NewSession(nil)
	_, err := sess.DeleteFrom(confirmationCodesTable).
		Where(dbr.Eq("employee_id", ID)).
		Exec()

	return err
}
