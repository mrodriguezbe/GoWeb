package service

import (
	"context"
	"fmt"
	"goweb/Desafio-Cierre/internal/repository"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicketMap
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicketMap) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets that travels to a destination proportionally
func (s *ServiceTicketDefault) GetProportionByDestination(dest string) (total int, err error) {
	total_amount, err := s.rp.Get(context.Background())

	if err != nil {
		fmt.Println("Service: ", err)
		return 0, err
	}

	total_dest, err := s.rp.GetTicketsByDestinationCountry(context.Background(), dest)

	if err != nil {
		fmt.Println("Service: ", err)
		return 0, err
	}

	return (len(total_dest) * 100) / len(total_amount), err
}

// GetTotalTickets returns the total number of tickets that travel to a destination
func (s *ServiceTicketDefault) GetTicketsByDestination(dest string) (total int, err error) {
	total_dest, err := s.rp.GetTicketsByDestinationCountry(context.Background(), dest)

	if err != nil {
		fmt.Println("Service: ", err)
		return 0, err
	}

	return len(total_dest), err
}
