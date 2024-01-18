package service

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
)

type Service struct {
	id          uuid.UUID
	description string
	price       float64
}

func New(description string, price float64) (*Service, error) {
	s := &Service{
		id: uuid.New(),
	}

	err := s.ChangeDescription(description)
	if err != nil {
		return nil, err
	}

	err = s.ChangePrice(price)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func Load(id string, description string, price float64) (*Service, error) {
	srvce, err := New(description, price)
	if err != nil {
		return nil, err
	}

	err = srvce.ChangeId(id)

	return srvce, err
}

func (s *Service) Description() string {
	return s.description
}

func (s *Service) Price() float64 {
	return s.price
}

func (s *Service) Id() uuid.UUID {
	return s.id
}

func (s *Service) ChangeId(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	serviceId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change service id")
	}

	s.id = serviceId

	return nil
}

func (s *Service) ChangeDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be empty")
	}

	s.description = description

	return nil
}

func (s *Service) ChangePrice(price float64) error {
	if price == 0 {
		return errors.New("price cannot be empty")
	}

	s.price = price

	return nil
}

func (s *Service) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id          string  `json:"id"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}{
		Id:          s.Id().String(),
		Description: s.Description(),
		Price:       s.Price(),
	})
}
