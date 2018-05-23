package password

// Passwords MUST
//
// Have an 8 character minimum
// AT LEAST 64 character maximum
// Allow all ASCII characters and spaces (unicode optional)
// Not be a common password (https://github.com/danielmiessler/SecLists/raw/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt)

import (
	"errors"
	"unicode"
)

var (
	// ErrTooShort is the error when a password doesn't meet the minimum password length requirement
	ErrTooShort = errors.New("Too Short")
	// ErrTooLong is the error when a password doesn't meet the maxumum password length requirement
	ErrTooLong = errors.New("Too Long")
	// ErrInvalidCharacters is the error when the password contains non-ASCII characters
	ErrInvalidCharacters = errors.New("Invalid Characters")
	// ErrCommon is the error when the password matches a string in the common password list
	ErrCommon = errors.New("Common Password")
)

const (
	maxCharacters = 64
	minCharacters = 8
)

// Password is a type alias for a string of password that needs to be validated
type Password string

// CheckValidity checks the Password based on the NIST password criteria
func (p *Password) CheckValidity(common CommonList) error {
	if !p.IsASCII() {
		return ErrInvalidCharacters
	}
	if p.UnderMinCharacters() {
		return ErrTooShort
	}
	if p.OverMaxCharacters() {
		return ErrTooLong
	}
	if p.IsCommon(common) {
		return ErrCommon
	}
	return nil
}

// IsASCII checks each rune in the string to determine if the string contains any non-ASCII characters
func (p *Password) IsASCII() bool {
	for _, r := range *p {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// OverMaxCharacters determines if the password has more than the configured maximum number of chatacters
func (p *Password) OverMaxCharacters() bool {
	return len(*p) >= maxCharacters
}

// UnderMinCharacters determines if the password has less than the configured minimum number of chatacters
func (p *Password) UnderMinCharacters() bool {
	return len(*p) < minCharacters
}

// IsCommon determines if the password is in the list of loaded common passwords
func (p *Password) IsCommon(common CommonList) bool {
	// TODO: This search should be more efficient
	found := common.Matches(string(*p))
	return found
}
