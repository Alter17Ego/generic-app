package identities

import (
	"strings"
)

type Identity int

const (
	Unknown Identity = iota
	Ktp
	DrivingLicense
	Kitas
	Passport
)

func ValidateIdentity(id Identity) bool {
	return id >= Unknown && id <= Passport
}

func FromString(identityStr string) Identity {
	str := strings.ToLower(identityStr)
	switch str {
	default:
		return Unknown
	case "ktp":
		return Ktp
	case "drivinglicense":
		return DrivingLicense
	case "kitas":
		return Kitas
	case "passport":
		return Passport
	}
}
