package location

import (
	"testing"

	"github.com/brother14th/locationmapping/db"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestGetLocationReport(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedLocationReport := db.LocationReport{gorm.Model{ID: 2}, "location2", 1234, 2000}
	mockObj := db.NewMockLocationRepository(mockCtrl)
	mockObj.EXPECT().GetLocationReport("location2").Return(expectedLocationReport, nil)
	userService := NewService(mockObj)
	locationReport, _ := userService.GetLocationReport("location2")

	if (locationReport.ID != expectedLocationReport.ID) || (locationReport.Location != expectedLocationReport.Location) ||
		(locationReport.TotalSquareFeet != expectedLocationReport.TotalSquareFeet) || (locationReport.PricePerMonth != expectedLocationReport.PricePerMonth) {
		t.Fatalf(`GetLocationReport("location2") =%v %q %v %v, expected %v %q %v %v`,
			locationReport.ID, locationReport.Location, locationReport.TotalSquareFeet, locationReport.PricePerMonth,
			expectedLocationReport.ID, locationReport.Location, expectedLocationReport.TotalSquareFeet, expectedLocationReport.PricePerMonth)
	}
}
