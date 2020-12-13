package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/brother14th/locationmapping/db"
)

func TestGetPreferredLocation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	expectedLocation := "location1"
	mockObj := db.NewMockUserRepository(mockCtrl)
	mockObj.EXPECT().GetPreferredLocation("user1").Return(expectedLocation, nil)
	userService := NewService(mockObj)
	location, _ := userService.GetPreferredLocation("user1")

	if location != expectedLocation {
		t.Fatalf(`PreferredLocation("user1") = %q, expected %q`, location, expectedLocation)
	}
}
