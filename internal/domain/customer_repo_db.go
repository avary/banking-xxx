package domain

import (
	"context"
	"database/sql"
	"github.com/ashtishad/banking/pkg/lib"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	findByIdSql = `SELECT id, name, date_of_birth,city,zipcode,status FROM customers WHERE id=$1;`
)

type CustomerRepoDb struct {
	db *pgxpool.Pool
}

func NewCustomerRepoDb(db *pgxpool.Pool) *CustomerRepoDb {
	return &CustomerRepoDb{db: db}
}

// FindById returns a customer by id
// Returns error if customer not found
func (d *CustomerRepoDb) FindById(id string) (*Customer, lib.RestErr) {
	// Note: Select * would supply data on db table order, order would mismatch with struct fields
	findByIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"
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
