package db

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//User maps to db row
type User struct {
	gorm.Model
	Username          string
	Password          string
	PreferredLocation string
}

//UserRepository provides access to db
type UserRepository interface {
	ValidateUser(username string, password string) (bool, error)
	SetPreferredLocation(username string, location string) (bool, error)
	GetPreferredLocation(username string) (string, error)
}

type userRepository struct {
	dbContext *gorm.DB
}

func (r *userRepository) ValidateUser(username string, password string) (bool, error) {
	var user User
	result := r.dbContext.First(&user, "username = ? and password = ?", username, password)
	if result.Error == nil && result.RowsAffected == 1 {
		return true, nil
	}
	return false, result.Error

}

func (r *userRepository) SetPreferredLocation(username string, location string) (bool, error) {
	var user User
	result := r.dbContext.First(&user, "username = ?", username)
	if result.Error == nil && result.RowsAffected == 1 {
		var loc Location
		// check if location exists in location table
		result = r.dbContext.First(&loc, "name = ?", location)
		if result.Error == nil && result.RowsAffected == 1 {
			result = r.dbContext.Model(&user).Update("preferred_location", location)
			if result.Error == nil && result.RowsAffected == 1 {
				return true, nil
			}
			return false, result.Error
		}
		return false, result.Error
	}
	return false, result.Error
}
func (r *userRepository) GetPreferredLocation(username string) (string, error) {
	var user User
	result := r.dbContext.First(&user, "username = ?", username)

	if result.Error == nil && result.RowsAffected == 1 {
		return user.PreferredLocation, nil
	}
	return "", result.Error

}

//NewUserRepository creates new repository
func NewUserRepository() (UserRepository, error) {
	dbPath := filepath.ToSlash(os.Getenv("GOPATH")) + "/src/github.com/brother14th/locationmapping/db/locationmapping.db"
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return &userRepository{
		dbContext: db,
	}, nil
}
