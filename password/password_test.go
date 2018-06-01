package password

import "testing"

func TestPasswordCheckValidity(t *testing.T) {
	var testStrings = []struct {
		value    Password
		expected error
	}{
		{Password(""), ErrTooShort},
		{Password("abcd"), ErrTooShort},
		{Password("qwertyui"), nil},
		{Password("qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopas"), nil},
		{Password("asdlfdja;sfjajfpojp8jp8efpw3jpowefjohefpowahepfo"), nil},
		{Password("asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo"), nil},
		{Password("asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo"), nil},
		{Password("asdlfdja;lsdijfa;sdfpoaisdufpoaisuf9oarpfauhrfiuehfpiudhfioufgoiudfhgoiudfgpdupodsifpuosiUFPAOSIDUFPAOSDIUFP"), ErrTooLong},
		{Password("☺☻☹"), ErrInvalidCharacters},
		{Password("日a本b語ç日ð本Ê語þ日¥本¼語i日©"), ErrInvalidCharacters},
		{Password("日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ"), ErrInvalidCharacters},
		{Password("\x80\x80\x80\x80asdasdaslkjdalskjf"), ErrInvalidCharacters},
		{Password("ಠ_ಠasdsadಠ_ಠ"), ErrInvalidCharacters},
		{Password("zzzzfitt"), ErrCommon},
		{Password("vjht008"), ErrTooShort},
		{Password("123456789"), ErrCommon},
	}
	filename := "password_list_test.txt"
	common, err := NewCommonList(filename)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range testStrings {
		got := test.value.CheckValidity(common)
		if got != test.expected {
			t.Errorf("%s -- expected: %v, got: %v", test.value, test.expected, got)
		}
	}
}
