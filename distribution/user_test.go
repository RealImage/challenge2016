package distribution_test

import (
	"testing"

	"github.com/binkkatal/challenge2016/distribution"
	"github.com/stretchr/testify/assert"
)

func init() {
	distribution.UniversalAreaList = distribution.RetrieveAreas("../cities.csv")
}

func TestValidUser(t *testing.T) {
	u, err := distribution.GetUser("1")
	assert.NoError(t, err, "Should not get any error")
	assert.Equal(t, "1", u.ID, "Should get correct user ID")
}

func TestInvalidUser(t *testing.T) {
	u, err := distribution.GetUser("5")
	assert.Error(t, err, "Should get error")
	assert.NotEqual(t, "5", u.ID, "User should be blank")
}

func TestUserWithInvalidaPermissions(t *testing.T) {
	u, err := distribution.GetUser("3")
	assert.NoError(t, err, "Should return a user")
	assert.Equal(t, "3", u.ID, "Should get correct user ID")

	userPermissions, permissionErr := u.ParsePermission()
	assert.Error(t, permissionErr, "Should return an error")
	assert.Equal(t, permissionErr.Error(), "Invalid Parent (2) for user (3)", "Should get correct description of error")
	assert.Nil(t, userPermissions, "Should not return any permissions for invalid declared permissions")
}

func TestPermissions(t *testing.T) {
	u, err := distribution.GetUser("2")
	assert.NoError(t, err, "Should not get error while retirieving a valid user")
	userPermissions, err := u.ParsePermission()
	assert.True(t, userPermissions["pathankot-punjab-india"], "User should be allowed for distribution in PATHANKOT-PUNJAB-INDIA ")
	assert.False(t, userPermissions["karnataka-india"], "User should be allowed for distribution in KARNATAKA-INDIA ")
	assert.False(t, userPermissions["yadgiri-karnataka-india"], "User should be allowed for distribution in YADGIRI-KARNATAKA-INDIA ")

}
