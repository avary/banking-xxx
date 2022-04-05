package domain

import (
	"github.com/ashtishad/banking/internal/dto"
	"github.com/ashtishad/banking/pkg/lib"
	"time"
)

type Customer struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Zipcode     string    `json:"zipcode"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Status      int8      `json:"status"`
}

// CustomerRepository is a SECONDARY PORT on Hexagonal architecture
type CustomerRepository interface {
	FindById(id int64) (*Customer, lib.RestErr)

	//FindByStatus(status string) ([]Customer, lib.RestErr)
	//FindByName(name string) (*Customer, lib.RestErr)
	//Create(customer *Customer) (*Customer, lib.RestErr)
	//Update(customer *Customer) (*Customer, lib.RestErr)
	//Delete(id string) lib.RestErr
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
