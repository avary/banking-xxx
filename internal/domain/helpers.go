package domain

import (
	"context"
	"database/sql"
	"github.com/ashtishad/banking/pkg/lib"
)

// checkCustomerExists checks if a customer exists in the database
func (d *CustomerRepoDb) checkCustomerExists(id int64) (bool, lib.RestErr) {
	row := d.db.QueryRow(context.Background(), findByIdSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, lib.NewInternalServerError("error when trying to get customer", err)
	}
}
