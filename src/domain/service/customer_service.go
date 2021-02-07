package service

import (
	"github.com/igson/banking/src/domain/dto"
	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

type customerService struct {
	repo interfaces.ICustomerRepository
}

//NewCustomerService acesso ao repositório
func NewCustomerService(customerRepository interfaces.ICustomerRepository) interfaces.ICustomerService {
	return &customerService{
		repo: customerRepository,
	}
}

func (s *customerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errors.RestErroAPI) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s *customerService) GetCustomer(id int64) (*dto.CustomerResponse, *errors.RestErroAPI) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}
