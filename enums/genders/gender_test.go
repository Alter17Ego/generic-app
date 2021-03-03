package genders_test

import (
	"testing"

	"github.com/Alter17Ego/generic-app/enums/genders"

	"github.com/stretchr/testify/assert"
)

func TestShouldConvertStrToGender(t *testing.T) {
	testArgs := map[string]genders.Gender{
		"Male":   genders.Male,
		"male":   genders.Male,
		"MALE":   genders.Male,
		"Female": genders.Female,
		"female": genders.Female,
		"FEMALE": genders.Female,
		"Other":  genders.Other,
		"other":  genders.Other,
		"OTHER":  genders.Other,
	}
	for key := range testArgs {
		testValue := key
		expectedResult := testArgs[key]
		assert.Equal(t, expectedResult, genders.FromString(testValue))
	}
}

func TestShouldConvertNotInvalidValueToUnkown(t *testing.T) {
	testArgs := map[string]genders.Gender{
		"mal":   genders.Unknown,
		"laki":  genders.Unknown,
		"pri":   genders.Unknown,
		"nest":  genders.Unknown,
		"perem": genders.Unknown,
		"man":   genders.Unknown,
	}
	for key := range testArgs {
		testValue := genders.FromString(key)
		expectedResult := testArgs[key]
		assert.Equal(t, expectedResult, testValue)
	}
}

func TestValueShouldBeValidGender(t *testing.T) {
	gndrs := []genders.Gender{genders.Male, genders.Female, genders.Other, genders.Unknown}
	for _, gender := range gndrs {
		isValidId := genders.ValidateGender(gender)
		assert.True(t, isValidId)
	}
}

func TestValueShouldBeInvalidGender(t *testing.T) {
	gndrs := []genders.Gender{5, 6, 7, 8, 9}
	for _, gender := range gndrs {
		isValidId := genders.ValidateGender(gender)
		assert.False(t, isValidId)
	}
}
