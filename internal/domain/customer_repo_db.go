package domain

import (
	"context"
	"database/sql"
	"github.com/ashtishad/banking/pkg/lib"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	findByIdSql = `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = $1;`
)

type CustomerRepoDb struct {
	db *pgxpool.Pool
}

func NewCustomerRepoDb(db *pgxpool.Pool) *CustomerRepoDb {
	return &CustomerRepoDb{db: db}
}

// FindById returns a customer by id
// Returns error if customer not found
func (d *CustomerRepoDb) FindById(id int64) (*Customer, lib.RestErr) {
	row := d.db.QueryRow(context.Background(), findByIdSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

	switch err {
	case sql.ErrNoRows:
		return nil, lib.NewNotFoundError("no customer found with given id")
	case nil:
		return &c, nil
	default:
		return nil, lib.NewInternalServerError("error when trying to get customer", err)
	}
}
