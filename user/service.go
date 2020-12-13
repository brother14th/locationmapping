package user

import (
	"errors"

	"github.com/brother14th/locationmapping/db"
)

//Service interface
type Service interface {
	SetPreferredLocation(string, string) (bool, error)
	GetPreferredLocation(string) (string, error)
}

type service struct {
	userRepository db.UserRepository
}

func (s *service) SetPreferredLocation(username string, location string) (bool, error) {

	if username == "" {
		return false, errors.New("empty username")
	}
	if location == "" {
		return false, errors.New("empty location")
	}
	ok, _ := s.userRepository.SetPreferredLocation(username, location)
	if !ok {
		return false, errors.New("error setting preferred location")
	}
	return true, nil
}

func (s *service) GetPreferredLocation(username string) (string, error) {
	if username == "" {
		return "", errors.New("empty username")
	}
	location, _ := s.userRepository.GetPreferredLocation(username)
	return location, nil
}

//NewService creates a new service
func NewService(userRepository db.UserRepository) Service {
	return &service{
		userRepository: userRepository,
	}
}
