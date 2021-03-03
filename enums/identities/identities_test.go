package identities_test

import (
	"testing"

	"github.com/Alter17Ego/generic-app/enums/identities"

	"github.com/stretchr/testify/assert"
)

func TestShouldConvertStrToIdentity(t *testing.T) {
	testArgs := map[string]identities.Identity{"ktp": identities.Ktp,
		"KTP":            identities.Ktp,
		"Ktp":            identities.Ktp,
		"kitas":          identities.Kitas,
		"KITAS":          identities.Kitas,
		"Kitas":          identities.Kitas,
		"drivinglicense": identities.DrivingLicense,
		"Drivinglicense": identities.DrivingLicense,
		"DRIVINGLICENSE": identities.DrivingLicense,
		"passport":       identities.Passport,
		"PASSPORT":       identities.Passport,
		"Passport":       identities.Passport,
		"kta":            identities.Unknown,
		"juni":           identities.Unknown,
	}
	for key := range testArgs {
		testValue := identities.FromString(key)
		expectedResult := testArgs[key]
		assert.Equal(t, expectedResult, testValue)
	}
}

func TestShouldConvertInvalidStrToUnknown(t *testing.T) {
	testArgs := map[string]identities.Identity{
		"kta":  identities.Unknown,
		"juni": identities.Unknown,
	}
	for key := range testArgs {
		testValue := identities.FromString(key)
		expectedResult := testArgs[key]
		assert.Equal(t, expectedResult, testValue)
	}
}

func TestValueShouldBeValidIdentity(t *testing.T) {
	ids := []identities.Identity{identities.DrivingLicense, identities.Kitas, identities.Ktp, identities.Passport, identities.Unknown}
	for _, id := range ids {
		isValidId := identities.ValidateIdentity(id)
		assert.True(t, isValidId)
	}
}

func TestValueShouldBeInvalidIdentity(t *testing.T) {
	ids := []identities.Identity{5, 6, 7, 8}
	for _, id := range ids {
		isValidId := identities.ValidateIdentity(id)
		assert.False(t, isValidId)
	}
}
