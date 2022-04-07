package service

import (
	"github.com/ashtishad/banking/internal/domain"
	"github.com/ashtishad/banking/internal/dto"
	"github.com/ashtishad/banking/pkg/lib"
)

// CustomerService is our PRIMARY PORT
type CustomerService interface {
	GetCustomerById(id int64) (*dto.CustomerResponse, lib.RestErr)
	SearchByStatus(status int8) ([]dto.CustomerResponse, lib.RestErr)
	CreateCustomer(req dto.NewCustomerRequest) (*dto.CustomerResponse, lib.RestErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

// GetCustomerById returns customer by id
func (s DefaultCustomerService) GetCustomerById(id int64) (*dto.CustomerResponse, lib.RestErr) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToCustomerResponse()

	return &resp, nil
}

// SearchByStatus returns customer by status
func (s DefaultCustomerService) SearchByStatus(status int8) ([]dto.CustomerResponse, lib.RestErr) {
	customers, err := s.repo.FindByStatus(status)
	if err != nil || len(customers) == 0 {
		return nil, err
	}

	resp := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		resp = append(resp, c.ToCustomerResponse())
	}
	return resp, err
}

func (s DefaultCustomerService) CreateCustomer(req dto.NewCustomerRequest) (*dto.CustomerResponse, lib.RestErr) {
	c := domain.ToNewCustomer(req)

	customer, err := s.repo.Create(c)
	if err != nil {
		return nil, err
	}

	resp := customer.ToCustomerResponse()
	return &resp, nil
}
