package genders

import "strings"

type Gender int

const (
	Unknown Gender = iota
	Male
	Female
	Other
)

func ValidateGender(gender Gender) bool {
	return gender >= Unknown && gender <= Other
}

func FromString(genderStr string) Gender {
	str := strings.ToLower(genderStr)
	switch str {
	default:
		return Unknown
	case "male":
		return Male
	case "female":
		return Female
	case "other":
		return Other
	}
}
