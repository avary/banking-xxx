package domain

import (
	"github.com/ashtishad/banking/internal/dto"
	"github.com/ashtishad/banking/pkg/lib"
	"time"
)

// Customer is a domain entity
// Id is int64 matched pg big serial
// DateOfBirth is time.Time and is matched pg timestamp
type Customer struct {
	Id          int64     `json:"id"`
	Name        string    `json:"full_name" binding:"required"`
	City        string    `json:"city" binding:"required"`
	Zipcode     string    `json:"zipcode" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	Status      int8      `json:"status"`
}

// CustomerRepository is a SECONDARY PORT on Hexagonal architecture
type CustomerRepository interface {
	FindById(id int64) (*Customer, lib.RestErr)
	FindByStatus(status int8) ([]Customer, lib.RestErr)
	Create(c Customer) (*Customer, lib.RestErr)
	Update(c Customer) (*Customer, lib.RestErr)

	//FindByName(name string) (*Customer, lib.RestErr)
	//Delete(id int64) lib.RestErr
}

// statusAsText returns 1= active, 0= inactive
func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == 0 {
		statusAsText = "inactive"
	}
	return statusAsText
}

// timeAsText returns date in format yyyy-mm-dd
func (c Customer) dateAsText() string {
	return c.DateOfBirth.Format(lib.DbDateLayout)
}

// ToCustomerResponse  converts Customer domain entity to DTO response
// there is some type conversion happening when I show it to user in response
// Like: Customer.DateOfBirth is time.Time, However on user level its string
func (c Customer) ToCustomerResponse() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.dateAsText(),
		Status:      c.statusAsText(),
	}
}
