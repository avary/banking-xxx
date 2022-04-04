package domain

import (
	"github.com/ashtishad/banking/internal/dto"
	"github.com/ashtishad/banking/internal/lib"
)

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

// CustomerRepository is a SECONDARY PORT on Hexagonal architecture
type CustomerRepository interface {
	FindById(id string) (*Customer, lib.RestErr)

	//FindByStatus(status string) ([]Customer, lib.RestErr)
	//FindByName(name string) (*Customer, lib.RestErr)
	//Create(customer *Customer) (*Customer, lib.RestErr)
	//Update(customer *Customer) (*Customer, lib.RestErr)
	//Delete(id string) lib.RestErr
}

// statusAsText returns 1= active, 0= inactive
func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToCustomerResponse() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}
