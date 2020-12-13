package location

import (
	"errors"

	"github.com/brother14th/locationmapping/db"
)

//Service provides location summary
type Service interface {
	GetLocationReport(string) (db.LocationReport, error)
}

type service struct {
	locationRepository db.LocationRepository
}

func (s *service) GetLocationReport(location string) (db.LocationReport, error) {

	if location == "" {
		return db.LocationReport{}, errors.New("empty location")
	}

	locationReport, err := s.locationRepository.GetLocationReport(location)
	//fmt.Println(locationReport)

	if err != nil {
		return db.LocationReport{}, errors.New("error getting location report")
	}
	return locationReport, nil
}

//NewService creates a new service
func NewService(locationRepository db.LocationRepository) Service {
	return &service{
		locationRepository: locationRepository,
	}
}
