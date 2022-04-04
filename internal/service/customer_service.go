package service

import (
	"github.com/ashtishad/banking/internal/domain"
	"github.com/ashtishad/banking/internal/dto"
	"github.com/ashtishad/banking/internal/lib"
)

// CustomerService is our PRIMARY PORT
type CustomerService interface {
	GetCustomerById(status string) ([]dto.CustomerResponse, lib.RestErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

// GetCustomerById returns customer by id
func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, lib.RestErr) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToCustomerResponse()

	return &resp, nil
}
