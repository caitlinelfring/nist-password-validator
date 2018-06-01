// Package password provides utilities to validate a password or list of passwords against a set of criteria
package password

// Passwords MUST
//
// Have an 8 character minimum
// AT LEAST 64 character maximum
// Allow all ASCII characters and spaces (unicode optional)
// Not be a common password (https://github.com/danielmiessler/SecLists/raw/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt)

import (
	"errors"
	"io"
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

// Validator is a password validator with configurable settings
type Validator struct {
	AcceptASCIIOnly    bool
	MaxCharacters      int
	MinCharacters      int
	commonPasswordList *CommonList
}

// NewValidator returns a password validator based on the configuration values supplied
func NewValidator(acceptASCIIOnly bool, minChar, maxChar int) *Validator {
	// TODO: Make sure minChar, maxChar are greater than 0 and make sure maxChar is greater than minChar
	v := Validator{
		AcceptASCIIOnly: acceptASCIIOnly,
		MaxCharacters:   maxChar,
		MinCharacters:   minChar,
	}
	return &v
}

// AddCommonPasswords constructs a CommonList from the supplied reader
func (v *Validator) AddCommonPasswords(r io.Reader) error {
	list, err := NewCommonList(r)
	v.commonPasswordList = &list
	return err
}

// ValidatePassword checks the Password based on the NIST password criteria
func (v *Validator) ValidatePassword(pass string) error {
	p := Password(pass)
	if v.AcceptASCIIOnly && !p.isASCII() {
		return ErrInvalidCharacters
	}
	if v.overMaxCharacters(p) {
		return ErrTooLong
	}
	if v.underMinCharacters(p) {
		return ErrTooShort
	}
	if v.isCommon(p) {
		return ErrCommon
	}
	return nil
}

// overMaxCharacters determines if the password has more than the configured maximum number of chatacters
func (v *Validator) overMaxCharacters(p Password) bool {
	return len(p) > v.MaxCharacters
}

// underMinCharacters determines if the password has less than the configured minimum number of chatacters
func (v *Validator) underMinCharacters(p Password) bool {
	return len(p) < v.MinCharacters
}

// isCommon determines if the password is in the list of loaded common passwords
func (v *Validator) isCommon(p Password) bool {
	return v.commonPasswordList.Matches(string(p))
}

// Password is a type alias for a string of password that needs to be validated
type Password string

// IsASCII checks each rune in the string to determine if the string contains any non-ASCII characters
func (p *Password) isASCII() bool {
	for _, r := range *p {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}
