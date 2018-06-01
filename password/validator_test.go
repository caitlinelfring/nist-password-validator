package password

import (
	"os"
	"testing"
)

func TestValidator(t *testing.T) {
	var testStrings = []struct {
		pass               string
		overMaxCharacters  bool
		underMinCharacters bool
		isCommon           bool
	}{
		{"", false, true, false},
		{"abcd", false, true, false},
		{"qwertyui", false, false, false},
		{"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopas", false, false, false},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowefjohefpowahepfo", false, false, false},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo", false, false, false},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo", false, false, false},
		{"asdlfdja;lsdijfa;sdfpoaisdufpoaisuf9oarpfauhrfiuehfpiudhfioufgoiudfhgoiudfgpdupodsifpuosiUFPAOSIDUFPAOSDIUFP", true, false, false},
		{"☺☻☹", false, false, false},
		{"日a本b語ç日ð本Ê語þ日¥本¼語i日©", false, false, false},
		{"日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ", true, false, false},
		{"\x80\x80\x80\x80asdasdaslkjdalskjf", false, false, false},
		{"ಠ_ಠasdsadಠ_ಠ", false, false, false},
		{"zzzzfitt", false, false, true},
		{"vjht008", false, true, true},
		{"123456789", false, false, true},
	}

	filename := "password_list_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: Add more Validator tests
	// acceptASCIIOnly = false
	// min > max
	// max = min = 0 (defaults?)
	validator := NewValidator(true, 8, 64)
	err = validator.AddCommonPasswords(file)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range testStrings {
		pass := Password(test.pass)
		isCommon := validator.isCommon(pass)
		if isCommon != test.isCommon {
			t.Errorf("%s (isCommon) -- expected: %v, got: %v", test.pass, test.isCommon, isCommon)
		}

		overMaxCharacters := validator.overMaxCharacters(pass)
		if overMaxCharacters != test.overMaxCharacters {
			t.Errorf("%s (overMaxCharacters) -- expected: %v, got: %v", test.pass, test.overMaxCharacters, overMaxCharacters)
		}

		underMinCharacters := validator.underMinCharacters(pass)
		if underMinCharacters != test.underMinCharacters {
			t.Errorf("%s (underMinCharacters) -- expected: %v, got: %v", test.pass, test.underMinCharacters, underMinCharacters)
		}
	}
}

func TestValidatorValidatePassword(t *testing.T) {
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
