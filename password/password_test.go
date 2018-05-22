package password

import (
	"fmt"
	"os"
	"testing"
)

var common CommonList

func init() {
	filename := "../10-million-password-list-top-1000000.txt"
	var err error
	common, err = NewCommonList(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestPasswordCheckValidity(t *testing.T) {
	var testStrings = []struct {
		value    Password
		expected error
	}{
		{Password(""), ErrTooShort},
		{Password("abcd"), ErrTooShort},
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
	}

	for _, test := range testStrings {
		got := test.value.CheckValidity(common)
		if got != test.expected {
			t.Errorf("%s -- expected: %v, got: %v", test.value, test.expected, got)
		}
	}
}
