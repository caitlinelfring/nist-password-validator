package password

import (
	"os"
	"testing"
)

func TestPasswordCheckValidity(t *testing.T) {
	var testStrings = []struct {
		value    string
		expected error
	}{
		{"", ErrTooShort},
		{"abcd", ErrTooShort},
		{"qwertyui", nil},
		{"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopas", nil},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowefjohefpowahepfo", nil},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo", nil},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo", nil},
		{"asdlfdja;lsdijfa;sdfpoaisdufpoaisuf9oarpfauhrfiuehfpiudhfioufgoiudfhgoiudfgpdupodsifpuosiUFPAOSIDUFPAOSDIUFP", ErrTooLong},
		{"☺☻☹", ErrInvalidCharacters},
		{"日a本b語ç日ð本Ê語þ日¥本¼語i日©", ErrInvalidCharacters},
		{"日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ", ErrInvalidCharacters},
		{"\x80\x80\x80\x80asdasdaslkjdalskjf", ErrInvalidCharacters},
		{"ಠ_ಠasdsadಠ_ಠ", ErrInvalidCharacters},
		{"zzzzfitt", ErrCommon},
		{"vjht008", ErrTooShort},
		{"123456789", ErrCommon},
	}
	filename := "password_list_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	validator := NewValidator(true, 8, 64)

	// Creating a new Validator and adding the common password list are two different steps,
	// so there's the potential that there will be no common passwords to check against.
	// Therefore, we should test for this scenario.
	for _, test := range testStrings {
		got := validator.ValidatePassword(test.value)
		// Since the common list hasn't been configured, if it's a common password,
		// the validator it shouldn't return an error
		if test.expected == ErrCommon {
			if err != nil {
				t.Errorf("%s -- expected: no error, got: %v", test.value, got)
			}
		} else if got != test.expected {
			t.Errorf("%s -- expected: %v, got: %v", test.value, test.expected, got)
		}
	}

	// Now lets add the common password list and run the same tests
	err = validator.AddCommonPasswords(file)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range testStrings {
		got := validator.ValidatePassword(test.value)
		if got != test.expected {
			t.Errorf("%s -- expected: %v, got: %v", test.value, test.expected, got)
		}
	}
}
