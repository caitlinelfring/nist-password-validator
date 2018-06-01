// Package password provides utilities to validate a password or list of passwords against a set of criteria
package password

// Passwords MUST
//
// Have an 8 character minimum
// AT LEAST 64 character maximum
// Allow all ASCII characters and spaces (unicode optional)
// Not be a common password (https://github.com/danielmiessler/SecLists/raw/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt)

import (
	"unicode"
)

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
