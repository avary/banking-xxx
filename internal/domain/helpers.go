package domain

import (
	"context"
	"database/sql"
	"github.com/ashtishad/banking/internal/dto"
	"github.com/ashtishad/banking/pkg/lib"
	"regexp"
	"strconv"
	"strings"
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

// ValidateNewCustomerRequest validates Customer domain entity
func ValidateNewCustomerRequest(c dto.NewCustomerRequest) lib.RestErr {
	rName := regexp.MustCompile(lib.NameRegex.Pattern)
	rCity := regexp.MustCompile(lib.CityRegex.Pattern)
	rZipcode := regexp.MustCompile(lib.ZipRegex.Pattern)
	rDOB := regexp.MustCompile(lib.DOBRegex.Pattern)

	if !rName.MatchString(c.Name) {
		return lib.NewBadRequestError(lib.NameRegex.Error)
	}
	if !rCity.MatchString(c.City) {
		return lib.NewBadRequestError(lib.CityRegex.Error)
	}
	if !rZipcode.MatchString(c.Zipcode) {
		return lib.NewBadRequestError(lib.ZipRegex.Error)
	}
	if !rDOB.MatchString(c.DateOfBirth) {
		return lib.NewBadRequestError(lib.DOBRegex.Error)
	}

	return nil
}

// ValidateUpdateCustomerRequest validates Customer domain entity
func ValidateUpdateCustomerRequest(req dto.CustomerUpdateRequest) lib.RestErr {
	rName := regexp.MustCompile(lib.NameRegex.Pattern)
	rCity := regexp.MustCompile(lib.CityRegex.Pattern)
	rZipcode := regexp.MustCompile(lib.ZipRegex.Pattern)

	if !rName.MatchString(req.Name) {
		return lib.NewBadRequestError(lib.NameRegex.Error)
	}
	if !rCity.MatchString(req.City) {
		return lib.NewBadRequestError(lib.CityRegex.Error)
	}
	if !rZipcode.MatchString(req.Zipcode) {
		return lib.NewBadRequestError(lib.ZipRegex.Error)
	}
	if req.Status != "0" && req.Status != "1" {
		return lib.NewBadRequestError("status must be 0 or 1")
	}

	return nil
}

// ToNewCustomer converts NewCustomerRequest to Customer domain entity
func ToNewCustomer(cr dto.NewCustomerRequest) Customer {
	return Customer{
		Name:        cr.Name,
		City:        cr.City,
		Zipcode:     cr.Zipcode,
		DateOfBirth: lib.DateAsTime(cr.DateOfBirth),
		Status:      1,
	}
}

// ToUpdateCustomer converts CustomerUpdateRequest to Customer domain entity
func ToUpdateCustomer(id int64, cr dto.CustomerUpdateRequest) Customer {
	status, _ := strconv.Atoi(cr.Status)
	return Customer{
		Id:      id,
		Name:    cr.Name,
		City:    cr.City,
		Zipcode: cr.Zipcode,
		Status:  int8(status),
	}
}

// setStatus sets status to 0 or 1
// it converts "active" to "0" and "inactive" to "1"
func setStatus(req dto.CustomerUpdateRequest) lib.RestErr {
	status := strings.ToLower(req.Status)
	switch status {
	case "0":
		req.Status = "0"
	case "1":
		req.Status = "1"
	case "active":
		req.Status = "1"
	case "inactive":
		req.Status = "0"
	default:
		return lib.NewBadRequestError("status is invalid")
	}
	return nil
}
