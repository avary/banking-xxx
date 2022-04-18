package domain

import (
	"context"
	"database/sql"
	"github.com/ashtishad/banking/pkg/lib"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	findByIdSql       = `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = $1;`
	findByStatusSql   = `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = $1;`
	createSql         = `INSERT INTO customers (name, city, zipcode, date_of_birth) VALUES ($1, $2, $3, $4) RETURNING customer_id;`
	findAllSql        = `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers;`
	updateCustomerSql = `UPDATE customers SET name=$1, city=$2, zipcode=$3 ,status=$4 WHERE customer_id=$5;`
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

// FindByStatus returns a paginated list of customers by status
// Returns error if no customers found or invalid status
func (d *CustomerRepoDb) FindByStatus(status int8) ([]Customer, lib.RestErr) {
	rows, err := d.db.Query(context.Background(), findByStatusSql, status)

	if err != nil {
		return nil, lib.NewInternalServerError("error when trying to get customers by status", err)
	}

	defer rows.Close()

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status); err != nil {
			return nil, lib.NewInternalServerError("error when trying to get customers", err)
		}
		customers = append(customers, c)
	}

	if len(customers) == 0 {
		return nil, lib.NewNotFoundError("no customers found")
	}

	return customers, nil
}

// Create creates a new customer
// ToDO: Why pointer to customer as parameter?
func (d *CustomerRepoDb) Create(c *Customer) (*Customer, lib.RestErr) {
	row := d.db.QueryRow(context.Background(), createSql, c.Name, c.City, c.Zipcode, c.DateOfBirth)

	if err := row.Scan(&c.Id); err != nil {
		return nil, lib.NewInternalServerError("error when trying to create customer", err)
	}

	return c, nil
}

// Update updatable fields of a customer
// fields that are not updatable are ignored(customer_id, date_of_birth)
func (d *CustomerRepoDb) Update(c Customer) (*Customer, lib.RestErr) {
	// check user exists
	if check, err := d.checkCustomerExists(c.Id); check == false {
		return nil, err
	}

	_, err := d.db.Exec(context.Background(), updateCustomerSql, c.Name, c.City, c.Zipcode, c.Status, c.Id)
	if err != nil {
		return nil, lib.NewInternalServerError("Could not update user : ", err)
	}
	// retrieve updated user
	return d.FindById(c.Id)
}
