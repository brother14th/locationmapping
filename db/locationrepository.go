package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Location maps to db row
type Location struct {
	gorm.Model
	name string
}

//LocationReport maps to db row
type LocationReport struct {
	gorm.Model
	Location        string  `gorm:"column:location"`
	TotalSquareFeet int     `gorm:"column:total_square_feet"`
	PricePerMonth   float32 `gorm:"column:price_per_month"`
}

//LocationRepository provides access to db
type LocationRepository interface {
	GetLocationReport(location string) (LocationReport, error)
}

type locationRepository struct {
	dbContext *gorm.DB
}

func (r *locationRepository) GetLocationReport(location string) (LocationReport, error) {
	var locReport LocationReport

	result := r.dbContext.First(&locReport, "location = ?", location)
	if result.Error == nil && result.RowsAffected == 1 {
		return locReport, nil
	}

	return locReport, result.Error
}

//NewLocationRepository creates new repository
func NewLocationRepository() (LocationRepository, error) {
	db, err := gorm.Open(sqlite.Open("C:/Users/hngkh/go/src/github.com/brother14th/locationmapping/db/locationmapping.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return &locationRepository{
		dbContext: db,
	}, nil
}
